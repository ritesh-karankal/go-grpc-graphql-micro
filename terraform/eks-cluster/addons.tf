# Core EKS Add-ons (managed by EKS so upgrades stay in sync with the cluster version)
resource "aws_eks_addon" "vpc_cni" {
  cluster_name                = aws_eks_cluster.eks.name
  addon_name                  = "vpc-cni"
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "OVERWRITE"
}

resource "aws_eks_addon" "kube_proxy" {
  cluster_name                = aws_eks_cluster.eks.name
  addon_name                  = "kube-proxy"
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "OVERWRITE"
}

# CoreDNS needs nodes to schedule onto
resource "aws_eks_addon" "coredns" {
  cluster_name                = aws_eks_cluster.eks.name
  addon_name                  = "coredns"
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "OVERWRITE"

  depends_on = [aws_eks_node_group.nodes]
}

# EBS CSI Driver: backs the Postgres PVCs. Also creates a default gp3 StorageClass,
# since EKS no longer marks any StorageClass as default.
resource "aws_eks_addon" "ebs_csi" {
  cluster_name                = aws_eks_cluster.eks.name
  addon_name                  = "aws-ebs-csi-driver"
  resolve_conflicts_on_create = "OVERWRITE"
  resolve_conflicts_on_update = "OVERWRITE"

  configuration_values = jsonencode({
    defaultStorageClass = {
      enabled = true
    }
  })

  depends_on = [aws_eks_node_group.nodes]
}
