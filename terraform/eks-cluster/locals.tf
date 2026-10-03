data "aws_caller_identity" "current" {}

data "aws_availability_zones" "available" {
  state = "available"
}

locals {
  cluster_name = "${var.project}-${var.environment}"
  account_id   = data.aws_caller_identity.current.account_id
  azs          = slice(data.aws_availability_zones.available.names, 0, var.az_count)

  # /24 public subnets at the bottom of the VPC, /20 private subnets for pods (VPC CNI uses VPC IPs)
  public_subnet_cidrs  = [for i in range(var.az_count) : cidrsubnet(var.vpc_cidr, 8, i)]
  private_subnet_cidrs = [for i in range(var.az_count) : cidrsubnet(var.vpc_cidr, 4, i + 1)]
  nat_gateway_count    = var.single_nat_gateway ? 1 : var.az_count

  # Bastion role is always admin; ARNs are built from the current account so no account ID lives in the repo.
  # Static map keys keep for_each plannable before the bastion role exists.
  cluster_admins = merge(
    { bastion = aws_iam_role.bastion.arn },
    { for u in var.admin_user_names : "user-${u}" => "arn:aws:iam::${local.account_id}:user/${u}" },
  )
}
