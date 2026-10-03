#!/bin/bash
# EC2 user data for the EKS bastion (Amazon Linux 2023). Rendered by Terraform templatefile().

set -euo pipefail

LOG_FILE="/var/log/bastion-user-data.log"

exec > >(tee -a "$LOG_FILE") 2>&1

echo "========================================"
echo "Bastion setup started: $(date)"
echo "========================================"

# --------------------------------------------
# 1. Base packages (AWS CLI and SSM agent ship with AL2023)
# --------------------------------------------
dnf install -y git jq tar bash-completion

# --------------------------------------------
# 2. kubectl matching the cluster minor version
# --------------------------------------------
KUBECTL_VERSION=$(curl -fsSL "https://dl.k8s.io/release/stable-${cluster_version}.txt")

curl -fsSLo /tmp/kubectl \
    "https://dl.k8s.io/release/$${KUBECTL_VERSION}/bin/linux/amd64/kubectl"

install -o root -g root -m 0755 /tmp/kubectl /usr/local/bin/kubectl
rm -f /tmp/kubectl

kubectl version --client

# --------------------------------------------
# 3. Helm
# --------------------------------------------
curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-4 | bash

helm version

# --------------------------------------------
# 4. Shared kubeconfig for every login user (ssm-user, ec2-user)
# --------------------------------------------
mkdir -p /etc/eks

aws eks update-kubeconfig \
    --region "${aws_region}" \
    --name "${cluster_name}" \
    --kubeconfig /etc/eks/kubeconfig

chmod 644 /etc/eks/kubeconfig

cat > /etc/profile.d/eks.sh <<'PROFILE'
export KUBECONFIG=/etc/eks/kubeconfig
alias k=kubectl
source <(kubectl completion bash)
complete -o default -F __start_kubectl k
PROFILE

echo "========================================"
echo "Bastion setup finished: $(date)"
echo "Cluster: ${cluster_name} (${aws_region})"
echo "========================================"
