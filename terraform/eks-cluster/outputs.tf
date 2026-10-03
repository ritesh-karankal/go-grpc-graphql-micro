output "aws_region" {
  description = "AWS region"
  value       = var.aws_region
}

output "cluster_name" {
  description = "EKS Cluster Name"
  value       = aws_eks_cluster.eks.name
}

output "cluster_endpoint" {
  description = "Endpoint for EKS Control Plane"
  value       = aws_eks_cluster.eks.endpoint
}

output "cluster_security_group_id" {
  description = "Security Group ID attached to EKS Cluster"
  value       = aws_eks_cluster.eks.vpc_config[0].cluster_security_group_id
}

output "configure_kubectl" {
  description = "Command to configure kubectl context"
  value       = "aws eks update-kubeconfig --region ${var.aws_region} --name ${aws_eks_cluster.eks.name}"
}
