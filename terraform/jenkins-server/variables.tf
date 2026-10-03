variable "aws_region" {
  description = "AWS Region to deploy the Jenkins server"
  type        = string
  default     = "eu-north-1"
}

variable "name" {
  description = "Name prefix for all Jenkins server resources"
  type        = string
  default     = "jenkins-server"
}

variable "vpc_cidr" {
  description = "CIDR block for the Jenkins VPC"
  type        = string
  default     = "10.1.0.0/16"
}

variable "public_subnet_cidr" {
  description = "CIDR block for the Jenkins public subnet"
  type        = string
  default     = "10.1.1.0/24"
}

variable "availability_zone" {
  description = "Availability Zone for the Jenkins server"
  type        = string
  default     = "eu-north-1a"
}

variable "instance_type" {
  description = "EC2 instance type (Jenkins + SonarQube + Docker builds need >= 8 GiB RAM)"
  type        = string
  default     = "t3.large"
}

variable "root_volume_size" {
  description = "Root EBS volume size in GiB"
  type        = number
  default     = 30
}

variable "admin_cidrs" {
  description = "CIDRs allowed to reach the SonarQube UI (port 9000)"
  type        = list(string)
}

variable "jenkins_cidrs" {
  description = "CIDRs allowed to reach the Jenkins UI (port 8080); must include GitHub hook IPs for webhooks"
  type        = list(string)
  default     = ["0.0.0.0/0"]
}
