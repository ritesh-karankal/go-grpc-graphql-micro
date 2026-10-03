# Database credentials in AWS Secrets Manager, synced into the cluster by External Secrets Operator.
# Passwords are generated as ephemeral values and written with write-only arguments,
# so they never appear in Terraform state, plan output or git.

locals {
  app_name     = "go-micro-shop"
  databases    = ["account", "order"]
  secrets_path = "${local.app_name}/${var.environment}"
}

ephemeral "random_password" "db" {
  for_each = toset(local.databases)

  length  = 32
  special = false # embedded in DATABASE_URL, keep it URL-safe
}

resource "aws_secretsmanager_secret" "db" {
  for_each = toset(local.databases)

  name                    = "${local.secrets_path}/${each.key}-db"
  description             = "Postgres credentials for the ${each.key} service (${var.environment})"
  recovery_window_in_days = var.secret_recovery_window_days
}

resource "aws_secretsmanager_secret_version" "db" {
  for_each = toset(local.databases)

  secret_id = aws_secretsmanager_secret.db[each.key].id
  secret_string_wo = jsonencode({
    POSTGRES_USER     = "postgres"
    POSTGRES_PASSWORD = ephemeral.random_password.db[each.key].result
    POSTGRES_DB       = "${each.key}_db"
  })

  # Written once. Bumping the version rotates the secret, but Postgres keeps the
  # password it was initialised with, so rotate the database user as well.
  secret_string_wo_version = 1
}

# IAM Role for External Secrets Operator, assumed through Pod Identity.
# Read-only, and only this environment's secrets.
resource "aws_iam_role" "external_secrets" {
  name = "${local.cluster_name}-external-secrets-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = ["sts:AssumeRole", "sts:TagSession"]
        Effect = "Allow"
        Principal = {
          Service = "pods.eks.amazonaws.com"
        }
      }
    ]
  })
}

resource "aws_iam_role_policy" "external_secrets" {
  name = "${local.cluster_name}-external-secrets-read"
  role = aws_iam_role.external_secrets.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "secretsmanager:GetSecretValue",
          "secretsmanager:DescribeSecret",
          "secretsmanager:ListSecretVersionIds",
        ]
        Effect   = "Allow"
        Resource = "arn:aws:secretsmanager:${var.aws_region}:${local.account_id}:secret:${local.secrets_path}/*"
      }
    ]
  })
}

resource "aws_eks_pod_identity_association" "external_secrets" {
  cluster_name    = aws_eks_cluster.eks.name
  namespace       = "external-secrets"
  service_account = "external-secrets"
  role_arn        = aws_iam_role.external_secrets.arn
}
