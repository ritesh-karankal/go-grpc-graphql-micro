# IAM Role for the Jenkins EC2 instance
resource "aws_iam_role" "jenkins_role" {
  name = "${var.name}-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      }
    ]
  })
}

# Attach AdministratorAccess: Jenkins runs Terraform that creates the EKS cluster,
# VPC and IAM roles, and pushes to ECR. Scope this down for anything beyond a lab.
resource "aws_iam_role_policy_attachment" "admin" {
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
  role       = aws_iam_role.jenkins_role.name
}

# Attach AmazonSSMManagedInstanceCore for Session Manager shell access
resource "aws_iam_role_policy_attachment" "ssm" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"
  role       = aws_iam_role.jenkins_role.name
}

# Instance Profile
resource "aws_iam_instance_profile" "jenkins_profile" {
  name = "${var.name}-profile"
  role = aws_iam_role.jenkins_role.name
}
