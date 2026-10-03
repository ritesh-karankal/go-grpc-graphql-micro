# Bastion host: the only admin path to the private EKS API.
# Lives in a private subnet with no public IP and no inbound rules; you reach it through SSM.

data "aws_ssm_parameter" "al2023_ami" {
  name = "/aws/service/ami-amazon-linux-latest/al2023-ami-kernel-default-x86_64"
}

# IAM Role for the bastion
resource "aws_iam_role" "bastion" {
  name = "${local.cluster_name}-bastion-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

# Attach AmazonSSMManagedInstanceCore for Session Manager access
resource "aws_iam_role_policy_attachment" "bastion_ssm" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  role       = aws_iam_role.bastion.name
}

# Allow the bastion to look up this cluster (needed by `aws eks update-kubeconfig`)
resource "aws_iam_role_policy" "bastion_eks" {
  name = "${local.cluster_name}-bastion-eks"
  role = aws_iam_role.bastion.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action   = ["eks:DescribeCluster", "eks:ListClusters"]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

# Allow the manual AWS Load Balancer Controller setup from the bastion (tutorial step):
# create the controller's IAM policy and role, and bind the role with Pod Identity.
# Attach is limited to that one policy, so the bastion can't grant itself or pods more.
resource "aws_iam_role_policy" "bastion_lb_controller_setup" {
  name = "${local.cluster_name}-bastion-lb-controller-setup"
  role = aws_iam_role.bastion.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid      = "ControllerPolicy"
        Action   = ["iam:CreatePolicy", "iam:GetPolicy", "iam:GetPolicyVersion"]
        Effect   = "Allow"
        Resource = "arn:aws:iam::${local.account_id}:policy/AWSLoadBalancerControllerIAMPolicy"
      },
      {
        Sid      = "ControllerRole"
        Action   = ["iam:CreateRole", "iam:GetRole", "iam:ListAttachedRolePolicies", "iam:PassRole"]
        Effect   = "Allow"
        Resource = "arn:aws:iam::${local.account_id}:role/${local.cluster_name}-lb-controller-role"
      },
      {
        Sid      = "AttachOnlyControllerPolicy"
        Action   = "iam:AttachRolePolicy"
        Effect   = "Allow"
        Resource = "arn:aws:iam::${local.account_id}:role/${local.cluster_name}-lb-controller-role"
        Condition = {
          ArnEquals = {
            "iam:PolicyARN" = "arn:aws:iam::${local.account_id}:policy/AWSLoadBalancerControllerIAMPolicy"
          }
        }
      },
      {
        Sid = "PodIdentity"
        Action = [
          "eks:CreatePodIdentityAssociation",
          "eks:ListPodIdentityAssociations",
          "eks:DescribePodIdentityAssociation",
        ]
        Effect   = "Allow"
        Resource = "*"
      }
    ]
  })
}

# Instance Profile
resource "aws_iam_instance_profile" "bastion" {
  name = "${local.cluster_name}-bastion-profile"
  role = aws_iam_role.bastion.name
}

# Security Group: no ingress at all, egress for SSM, package installs and the EKS API
resource "aws_security_group" "bastion" {
  name        = "${local.cluster_name}-bastion-sg"
  description = "EKS bastion (SSM only, no inbound)"
  vpc_id      = aws_vpc.eks_vpc.id

  egress {
    description = "All outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${local.cluster_name}-bastion-sg"
  }
}

# Let the bastion reach the private EKS API endpoint
resource "aws_vpc_security_group_ingress_rule" "cluster_api_from_bastion" {
  security_group_id            = aws_eks_cluster.eks.vpc_config[0].cluster_security_group_id
  referenced_security_group_id = aws_security_group.bastion.id
  from_port                    = 443
  to_port                      = 443
  ip_protocol                  = "tcp"
  description                  = "EKS API from bastion"
}

# Bastion EC2 Instance
resource "aws_instance" "bastion" {
  ami                         = data.aws_ssm_parameter.al2023_ami.value
  instance_type               = var.bastion_instance_type
  subnet_id                   = aws_subnet.private[0].id
  vpc_security_group_ids      = [aws_security_group.bastion.id]
  iam_instance_profile        = aws_iam_instance_profile.bastion.name
  associate_public_ip_address = false

  user_data = templatefile("${path.module}/bastion-setup.sh", {
    cluster_name    = aws_eks_cluster.eks.name
    cluster_version = var.cluster_version
    aws_region      = var.aws_region
  })
  user_data_replace_on_change = true

  metadata_options {
    http_tokens = "required" # IMDSv2 only
  }

  root_block_device {
    volume_size = 20
    volume_type = "gp3"
    encrypted   = true
  }

  tags = {
    Name = "${local.cluster_name}-bastion"
  }

  lifecycle {
    ignore_changes = [ami] # don't replace the bastion when a newer AMI is published
  }

  # NAT must exist so SSM agent and installs can reach the internet on first boot
  depends_on = [
    aws_route_table_association.private,
    aws_nat_gateway.nat,
  ]
}
