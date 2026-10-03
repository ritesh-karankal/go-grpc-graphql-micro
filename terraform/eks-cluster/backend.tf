# Remote state in S3 with native S3 locking. Bucket (contains the account ID)
# and per-environment key are passed at init time:
#   terraform init -backend-config=backend.hcl -backend-config="key=eks-cluster/dev/terraform.tfstate"
terraform {
  backend "s3" {
    region       = "eu-north-1"
    encrypt      = true
    use_lockfile = true
  }
}
