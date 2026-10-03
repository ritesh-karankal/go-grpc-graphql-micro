# dev: cheaper, single NAT, 2 AZs. Recreate freely.
environment        = "dev"
vpc_cidr           = "10.10.0.0/16"
az_count           = 2
single_nat_gateway = true
log_retention_days = 7

node_instance_types = ["t3.large"]
node_capacity_type  = "ON_DEMAND"
desired_size        = 2
min_size            = 2
max_size            = 4
