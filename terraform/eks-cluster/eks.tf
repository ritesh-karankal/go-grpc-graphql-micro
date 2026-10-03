# IAM Role for EKS Cluster
resource "aws_iam_role" "eks_cluster_role" {
  name = "${var.cluster_name}-cluster-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "eks.amazonaws.com"
        }
      }
    ]
  })
}

# Attach AmazonEKSClusterPolicy
resource "aws_iam_role_policy_attachment" "eks_cluster_policy" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
  role       = aws_iam_role.eks_cluster_role.name
}

# Attach AmazonEKSVPCResourceController
resource "aws_iam_role_policy_attachment" "eks_vpc_resource_controller" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"
  role       = aws_iam_role.eks_cluster_role.name
}

# EKS Cluster Definition
resource "aws_eks_cluster" "eks" {
  name     = var.cluster_name
  role_arn = aws_iam_role.eks_cluster_role.arn
  version  = var.cluster_version

  # Access entries instead of the aws-auth ConfigMap; admins are granted explicitly
  # below so access doesn't depend on who happened to run `terraform apply`.
  access_config {
    authentication_mode                         = "API"
    bootstrap_cluster_creator_admin_permissions = false
  }

  vpc_config {
    subnet_ids              = concat(aws_subnet.public[*].id, aws_subnet.private[*].id)
    endpoint_public_access  = true
    endpoint_private_access = true
  }

  depends_on = [
    aws_iam_role_policy_attachment.eks_cluster_policy,
    aws_iam_role_policy_attachment.eks_vpc_resource_controller,
  ]

  tags = {
    Name = var.cluster_name
  }
}

# Cluster-admin Access for Jenkins and your IAM user
# ARNs are built from the current account so no account ID lives in the repo.
data "aws_caller_identity" "current" {}

locals {
  cluster_admin_arns = concat(
    ["arn:aws:iam::${data.aws_caller_identity.current.account_id}:role/${var.jenkins_role_name}"],
    [for u in var.admin_user_names : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:user/${u}"],
  )
}

resource "aws_eks_access_entry" "admins" {
  for_each      = toset(local.cluster_admin_arns)
  cluster_name  = aws_eks_cluster.eks.name
  principal_arn = each.value
}

resource "aws_eks_access_policy_association" "admins" {
  for_each      = toset(local.cluster_admin_arns)
  cluster_name  = aws_eks_cluster.eks.name
  principal_arn = each.value
  policy_arn    = "arn:aws:eks::aws:cluster-access-policy/AmazonEKSClusterAdminPolicy"

  access_scope {
    type = "cluster"
  }

  depends_on = [aws_eks_access_entry.admins]
}
