variable "aws_region" {
  description = "AWS Region to deploy EKS cluster"
  type        = string
  default     = "eu-north-1"
}

variable "cluster_name" {
  description = "Name of the EKS Cluster"
  type        = string
  default     = "go-microservices-eks"
}

variable "cluster_version" {
  description = "Kubernetes version for the EKS Cluster"
  type        = string
  default     = "1.36"
}

variable "jenkins_role_name" {
  description = "IAM role of the Jenkins server, granted cluster-admin"
  type        = string
  default     = "jenkins-server-role"
}

variable "admin_user_names" {
  description = "IAM users granted cluster-admin (for kubectl from your laptop)"
  type        = list(string)
  default     = ["devops-user"]
}

variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  default     = "10.0.0.0/16"
}

variable "public_subnet_cidrs" {
  description = "CIDR blocks for public subnets"
  type        = list(string)
  default     = ["10.0.1.0/24", "10.0.2.0/24"]
}

variable "private_subnet_cidrs" {
  description = "CIDR blocks for private subnets"
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.20.0/24"]
}

variable "availability_zones" {
  description = "Availability Zones in eu-north-1"
  type        = list(string)
  default     = ["eu-north-1a", "eu-north-1b"]
}

variable "node_instance_type" {
  description = "EC2 Instance type for EKS worker nodes"
  type        = list(string)
  default     = ["t3.large"]
}

variable "desired_size" {
  description = "Desired number of worker nodes"
  type        = number
  default     = 2
}

variable "min_size" {
  description = "Minimum number of worker nodes"
  type        = number
  default     = 1
}

variable "max_size" {
  description = "Maximum number of worker nodes"
  type        = number
  default     = 4
}
