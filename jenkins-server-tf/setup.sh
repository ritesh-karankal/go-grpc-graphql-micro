#!/bin/bash

set -euo pipefail

LOG_FILE="/var/log/ec2-user-data.log"

exec > >(tee -a "$LOG_FILE") 2>&1

echo "========================================"
echo "EC2 User Data started: $(date)"
echo "========================================"

export DEBIAN_FRONTEND=noninteractive

# --------------------------------------------
# 1. System update + basic packages
# --------------------------------------------

echo "========================================"
echo "[1/10] Updating system..."
echo "========================================"

apt-get update -y

apt-get install -y \
    ca-certificates \
    curl \
    wget \
    gnupg \
    unzip \
    tar \
    fontconfig


# --------------------------------------------
# 2. Java 21
# --------------------------------------------

echo "========================================"
echo "[2/10] Installing Java 21..."
echo "========================================"

apt-get install -y openjdk-21-jre

java -version


# --------------------------------------------
# 3. Jenkins
# --------------------------------------------

echo "========================================"
echo "[3/10] Installing Jenkins..."
echo "========================================"

mkdir -p /etc/apt/keyrings

curl -fsSL \
    https://pkg.jenkins.io/debian-stable/jenkins.io-2026.key \
    -o /etc/apt/keyrings/jenkins-keyring.asc

cat > /etc/apt/sources.list.d/jenkins.list <<EOF
deb [signed-by=/etc/apt/keyrings/jenkins-keyring.asc] https://pkg.jenkins.io/debian-stable binary/
EOF

apt-get update -y

apt-get install -y jenkins

systemctl enable jenkins
systemctl start jenkins

systemctl status jenkins --no-pager || true


# --------------------------------------------
# 4. Docker
# --------------------------------------------
echo "========================================"
echo "[4/10] Installing Docker..."
echo "========================================"

install -m 0755 -d /etc/apt/keyrings

curl -fsSL \
    https://download.docker.com/linux/ubuntu/gpg \
    -o /etc/apt/keyrings/docker.asc

chmod a+r /etc/apt/keyrings/docker.asc

cat > /etc/apt/sources.list.d/docker.sources <<EOF
Types: deb
URIs: https://download.docker.com/linux/ubuntu
Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
Components: stable
Architectures: $(dpkg --print-architecture)
Signed-By: /etc/apt/keyrings/docker.asc
EOF

apt-get update -y

apt-get install -y \
    docker-ce \
    docker-ce-cli \
    containerd.io \
    docker-buildx-plugin \
    docker-compose-plugin

systemctl enable docker
systemctl start docker

# Allow Jenkins and Ubuntu users to run Docker
usermod -aG docker jenkins
usermod -aG docker ubuntu

systemctl restart docker

docker --version
docker compose version


# --------------------------------------------
# 5. SonarQube
# --------------------------------------------
echo "========================================"
echo "[5/10] Installing SonarQube..."
echo "========================================"

# Remove old container if it exists
docker rm -f sonar 2>/dev/null || true

docker run -d \
    --name sonar \
    --restart unless-stopped \
    -p 9000:9000 \
    sonarqube:lts-community

echo "SonarQube container started."


# --------------------------------------------
# 6. AWS CLI
# --------------------------------------------
echo "========================================"
echo "[6/10] Installing AWS CLI..."
echo "========================================"

cd /tmp

# Clean temporary files from any previous attempt
rm -rf /tmp/aws
rm -f /tmp/awscliv2.zip

# Download AWS CLI v2
curl -fsSL \
    "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" \
    -o /tmp/awscliv2.zip

# Extract
unzip -q /tmp/awscliv2.zip -d /tmp

# Install
/tmp/aws/install \
    --install-dir /usr/local/aws-cli \
    --bin-dir /usr/local/bin

# Clean temporary files
rm -rf /tmp/aws
rm -f /tmp/awscliv2.zip

# Verify
/usr/local/bin/aws --version


# --------------------------------------------
# 7. kubectl
# --------------------------------------------
echo "========================================"
echo "[7/10] Installing kubectl..."
echo "========================================"

KUBECTL_VERSION=$(curl -L -s https://dl.k8s.io/release/stable.txt)

curl -fsSLo /tmp/kubectl \
    "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl"

install -o root -g root -m 0755 \
    /tmp/kubectl \
    /usr/local/bin/kubectl

rm -f /tmp/kubectl

kubectl version --client


# --------------------------------------------
# 8. eksctl
# --------------------------------------------
echo "========================================"
echo "[8/10] Installing eksctl..."
echo "========================================"

ARCH=amd64
PLATFORM="$(uname -s)_${ARCH}"

cd /tmp

curl -fsSLO \
    "https://github.com/eksctl-io/eksctl/releases/latest/download/eksctl_${PLATFORM}.tar.gz"

tar -xzf "eksctl_${PLATFORM}.tar.gz"

install -m 0755 \
    /tmp/eksctl \
    /usr/local/bin/eksctl

rm -f \
    "/tmp/eksctl_${PLATFORM}.tar.gz" \
    /tmp/eksctl

eksctl version


# --------------------------------------------
# 9. Terraform
# --------------------------------------------
echo "========================================"
echo "[9/10] Installing Terraform..."
echo "========================================"

curl -fsSL \
    https://apt.releases.hashicorp.com/gpg \
    | gpg --dearmor \
    -o /usr/share/keyrings/hashicorp-archive-keyring.gpg

chmod 644 /usr/share/keyrings/hashicorp-archive-keyring.gpg

echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(grep -oP '(?<=UBUNTU_CODENAME=).*' /etc/os-release || lsb_release -cs) main" \
    > /etc/apt/sources.list.d/hashicorp.list

apt-get update -y

apt-get install -y terraform

terraform version


# --------------------------------------------
# 10. Trivy
# --------------------------------------------
echo "========================================"
echo "[10/10] Installing Trivy..."
echo "========================================"

curl -fsSL \
    https://aquasecurity.github.io/trivy-repo/deb/public.key \
    | gpg --dearmor \
    -o /usr/share/keyrings/trivy.gpg

chmod 644 /usr/share/keyrings/trivy.gpg

echo "deb [signed-by=/usr/share/keyrings/trivy.gpg] https://aquasecurity.github.io/trivy-repo/deb generic main" \
    > /etc/apt/sources.list.d/trivy.list

apt-get update -y

apt-get install -y trivy

trivy --version


# --------------------------------------------
# Helm
# --------------------------------------------
echo "========================================"
echo "Installing Helm..."
echo "========================================"

curl -fsSL https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-4 | bash

helm version


# --------------------------------------------
# Final verification
# --------------------------------------------

echo ""
echo "========================================"
echo "INSTALLATION COMPLETE"
echo "========================================"

echo "Java:"
java -version

echo ""
echo "Jenkins:"
systemctl is-active jenkins || true

echo ""
echo "Docker:"
docker --version

echo ""
echo "AWS CLI:"
aws --version

echo ""
echo "kubectl:"
kubectl version --client

echo ""
echo "eksctl:"
eksctl version

echo ""
echo "Terraform:"
terraform version

echo ""
echo "Trivy:"
trivy --version

echo ""
echo "Helm:"
helm version

echo ""
echo "SonarQube:"
docker ps --filter name=sonar

echo ""
echo "========================================"
echo "EC2 User Data finished: $(date)"
echo "Log: $LOG_FILE"
echo "========================================"