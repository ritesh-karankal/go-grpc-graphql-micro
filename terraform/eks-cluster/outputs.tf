output "aws_region" {
  description = "AWS region"
  value       = var.aws_region
}

output "environment" {
  description = "Deployment environment"
  value       = var.environment
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

output "vpc_id" {
  description = "EKS VPC ID"
  value       = aws_vpc.eks_vpc.id
}

output "public_subnet_ids" {
  description = "Public subnet IDs (for load balancers)"
  value       = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  description = "Private subnet IDs (for nodes)"
  value       = aws_subnet.private[*].id
}

output "bastion_instance_id" {
  description = "Bastion host instance ID"
  value       = aws_instance.bastion.id
}

output "bastion_ssm_session" {
  description = "Command to open a shell on the bastion (kubectl is preconfigured there)"
  value       = "aws ssm start-session --region ${var.aws_region} --target ${aws_instance.bastion.id}"
}
