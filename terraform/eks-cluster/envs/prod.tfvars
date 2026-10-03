# prod: 3 AZs with a NAT Gateway each, one node per AZ minimum.
environment        = "prod"
vpc_cidr           = "10.20.0.0/16"
az_count           = 3
single_nat_gateway = false
log_retention_days = 30

node_instance_types = ["t3.large"]
node_capacity_type  = "ON_DEMAND"
desired_size        = 3
min_size            = 3
max_size            = 6
