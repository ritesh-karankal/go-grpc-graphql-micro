# Remote state in S3 with native S3 locking (no DynamoDB table needed).
# The bucket name contains the AWS account ID, so it is passed at init time:
#   terraform init -backend-config=backend.hcl   (see backend.hcl.example)
terraform {
  backend "s3" {
    key          = "jenkins-server/terraform.tfstate"
    region       = "eu-north-1"
    encrypt      = true
    use_lockfile = true
  }
}
