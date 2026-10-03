variable "aws_region" {
  description = "AWS Region to deploy EKS cluster"
  type        = string
  default     = "eu-north-1"
}

variable "project" {
  description = "Project name, used as the resource name prefix"
  type        = string
  default     = "go-microservices"
}

variable "environment" {
  description = "Deployment environment (dev or prod)"
  type        = string

  validation {
    condition     = contains(["dev", "prod"], var.environment)
    error_message = "environment must be dev or prod."
  }
}

variable "cluster_version" {
  description = "Kubernetes version for the EKS Cluster"
  type        = string
  default     = "1.36"
}

# --------------------------------------------
# Network
# --------------------------------------------

variable "vpc_cidr" {
  description = "CIDR block for VPC (must not overlap other environments or the Jenkins VPC)"
  type        = string
}

variable "az_count" {
  description = "Number of Availability Zones to spread subnets and nodes across"
  type        = number
  default     = 3
}

variable "single_nat_gateway" {
  description = "Use one NAT Gateway for all AZs (cheaper, not AZ-fault-tolerant)"
  type        = bool
  default     = false
}

# --------------------------------------------
# Cluster access
# --------------------------------------------

variable "admin_user_names" {
  description = "Extra IAM users granted cluster-admin. Empty by default: the API is private and admin access goes through the bastion."
  type        = list(string)
  default     = []
}

variable "bastion_instance_type" {
  description = "EC2 instance type for the bastion host"
  type        = string
  default     = "t3.micro"
}

variable "log_retention_days" {
  description = "Retention for EKS control plane logs in CloudWatch"
  type        = number
  default     = 30
}

# --------------------------------------------
# Nodes
# --------------------------------------------

variable "node_instance_types" {
  description = "EC2 Instance types for EKS worker nodes"
  type        = list(string)
  default     = ["t3.large"]
}

variable "node_capacity_type" {
  description = "ON_DEMAND or SPOT"
  type        = string
  default     = "ON_DEMAND"
}

variable "node_disk_size" {
  description = "Root EBS volume size (GiB) for worker nodes"
  type        = number
  default     = 30
}

variable "desired_size" {
  description = "Desired number of worker nodes"
  type        = number
}

variable "min_size" {
  description = "Minimum number of worker nodes"
  type        = number
}

variable "max_size" {
  description = "Maximum number of worker nodes"
  type        = number
}
