# VPC Definition
resource "aws_vpc" "eks_vpc" {
  cidr_block           = var.vpc_cidr
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "${local.cluster_name}-vpc"
  }
}

# Internet Gateway for Public Subnets
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.eks_vpc.id

  tags = {
    Name = "${local.cluster_name}-igw"
  }
}

# Public Subnets (load balancers, NAT Gateways)
resource "aws_subnet" "public" {
  count             = var.az_count
  vpc_id            = aws_vpc.eks_vpc.id
  cidr_block        = local.public_subnet_cidrs[count.index]
  availability_zone = local.azs[count.index]

  tags = {
    Name                                          = "${local.cluster_name}-public-${local.azs[count.index]}"
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/elb"                      = "1"
  }
}

# Private Subnets (worker nodes and pods)
resource "aws_subnet" "private" {
  count             = var.az_count
  vpc_id            = aws_vpc.eks_vpc.id
  cidr_block        = local.private_subnet_cidrs[count.index]
  availability_zone = local.azs[count.index]

  tags = {
    Name                                          = "${local.cluster_name}-private-${local.azs[count.index]}"
    "kubernetes.io/cluster/${local.cluster_name}" = "shared"
    "kubernetes.io/role/internal-elb"             = "1"
  }
}

# Elastic IPs for NAT Gateways
resource "aws_eip" "nat" {
  count  = local.nat_gateway_count
  domain = "vpc"

  tags = {
    Name = "${local.cluster_name}-nat-eip-${count.index + 1}"
  }
}

# NAT Gateways: one per AZ in prod so losing an AZ doesn't cut egress for the others
resource "aws_nat_gateway" "nat" {
  count         = local.nat_gateway_count
  allocation_id = aws_eip.nat[count.index].id
  subnet_id     = aws_subnet.public[count.index].id

  tags = {
    Name = "${local.cluster_name}-nat-gw-${count.index + 1}"
  }

  depends_on = [aws_internet_gateway.igw]
}

# Route Table for Public Subnets
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.eks_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw.id
  }

  tags = {
    Name = "${local.cluster_name}-public-rt"
  }
}

# Route Tables for Private Subnets (one per AZ, pointing at that AZ's NAT or the single NAT)
resource "aws_route_table" "private" {
  count  = var.az_count
  vpc_id = aws_vpc.eks_vpc.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat[var.single_nat_gateway ? 0 : count.index].id
  }

  tags = {
    Name = "${local.cluster_name}-private-rt-${local.azs[count.index]}"
  }
}

# Associate Public Subnets with Public Route Table
resource "aws_route_table_association" "public" {
  count          = var.az_count
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

# Associate Private Subnets with their Private Route Table
resource "aws_route_table_association" "private" {
  count          = var.az_count
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private[count.index].id
}
