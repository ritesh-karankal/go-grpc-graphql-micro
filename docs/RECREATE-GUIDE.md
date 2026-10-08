# Recreate This Project – Step-by-Step Guide

**For:** a junior DevOps engineer who wants to rebuild this whole platform in their own AWS account and understand every command along the way.

**You will build:** Jenkins + SonarQube on EC2 → a private EKS cluster with a bastion → the AWS Load Balancer Controller → Argo CD → CI pipelines with security gates → GitOps deployment → secrets from AWS Secrets Manager → Prometheus and Grafana → OpenTelemetry tracing, metrics and logs in SigNoz.

**Time:** one long day, or two relaxed ones. **Cost:** roughly $10–12 per day while everything runs, so read [Step 18 – Teardown](#step-18--teardown) before you start.

> **How to read this guide.** Every step has: **Goal** (what you're doing), **Why** (the reason), **Do** (commands to run), **Check** (how to know it worked) and **If it fails** (the errors we actually hit). Placeholders look like `<THIS>`. Replace them with your own values.

---

## Contents

- [Step 0 – Prerequisites](#step-0--prerequisites)
- [Step 1 – Fork and prepare the repository](#step-1--fork-and-prepare-the-repository)
- [Step 2 – Create the Terraform state bucket](#step-2--create-the-terraform-state-bucket)
- [Step 3 – Create the Jenkins server](#step-3--create-the-jenkins-server)
- [Step 4 – First-time Jenkins and SonarQube setup](#step-4--first-time-jenkins-and-sonarqube-setup)
- [Step 5 – Create the EKS cluster from Jenkins](#step-5--create-the-eks-cluster-from-jenkins)
- [Step 6 – Use the bastion to reach the cluster](#step-6--use-the-bastion-to-reach-the-cluster)
- [Step 7 – Install the AWS Load Balancer Controller](#step-7--install-the-aws-load-balancer-controller)
- [Step 8 – Install Argo CD](#step-8--install-argo-cd)
- [Step 9 – Create the SonarQube projects](#step-9--create-the-sonarqube-projects)
- [Step 10 – Create the ECR repositories](#step-10--create-the-ecr-repositories)
- [Step 11 – Remaining Jenkins setup (tools and credentials)](#step-11--remaining-jenkins-setup-tools-and-credentials)
- [Step 12 – Install External Secrets Operator](#step-12--install-external-secrets-operator)
- [Step 13 – Run the CI pipelines](#step-13--run-the-ci-pipelines)
- [Step 14 – Deploy with Argo CD](#step-14--deploy-with-argo-cd)
- [Step 15 – Prometheus and Grafana](#step-15--prometheus-and-grafana)
- [Step 16 – SigNoz and OpenTelemetry](#step-16--signoz-and-opentelemetry)
- [Step 17 – Prove GitOps works](#step-17--prove-gitops-works)
- [Step 18 – Teardown](#step-18--teardown)
- [Appendix A – Glossary](#appendix-a--glossary)
- [Appendix B – Command cheat sheet](#appendix-b--command-cheat-sheet)
- [Appendix C – Troubleshooting index](#appendix-c--troubleshooting-index)

---

## Step 0 – Prerequisites

### Accounts
- An **AWS account** with an IAM user (not root) that has `AdministratorAccess`. We call it `devops-user`.
- A **GitHub account**.
- An **NVD API key** (free) from <https://nvd.nist.gov/developers/request-an-api-key>. It arrives by email in minutes; you need it in Step 11.

### Tools on your laptop (Linux or macOS)

| Tool | Check | Why |
|---|---|---|
| AWS CLI v2 | `aws --version` | Talk to AWS |
| Session Manager plugin | `session-manager-plugin --version` | `aws ssm start-session` (shell without SSH) |
| Terraform ≥ 1.11 | `terraform version` | Write-only arguments used for secrets |
| Git | `git --version` | |
| Docker (optional) | `docker --version` | Run the app locally |

Install the Session Manager plugin: <https://docs.aws.amazon.com/systems-manager/latest/userguide/session-manager-working-with-install-plugin.html>.

### Configure the AWS CLI

```bash
aws configure                    # access key, secret, region eu-north-1, output json
aws sts get-caller-identity      # shows your account ID and arn:...:user/devops-user
```

**Write down your account ID.** We'll call it `<ACCOUNT_ID>`.

### Find your public IP
```bash
curl https://checkip.amazonaws.com    # <YOUR_IP>
```
Some services will only accept connections from this IP. If your IP changes later (home internet does that), see [Appendix C](#appendix-c--troubleshooting-index).

---

## Step 1 – Fork and prepare the repository

**Goal:** have your own copy that Jenkins and Argo CD can read and that the pipelines can push to.

**Do:**
1. On GitHub, fork `ritesh-karankal/go-grpc-graphql-micro`, then clone your fork:
   ```bash
   git clone git@github.com:<YOUR_GH_USER>/go-grpc-graphql-micro.git
   cd go-grpc-graphql-micro
   ```
2. Replace the original repo owner in three places. They are hard-coded because the pipelines push back to the repo:
   ```bash
   grep -rn "ritesh-karankal" jenkins/ k8s/scripts/
   ```
   - `jenkins/Jenkinsfile-Backend` and `jenkins/Jenkinsfile-Frontend`: the `git push https://…github.com/<owner>/go-grpc-graphql-micro.git` URL and `git config user.name / user.email`.
   - `k8s/scripts/create-argocd-app.sh`: `REPO_URL`.
3. Reset the image tags to a placeholder (the values in the repo refer to someone else's builds):
   ```bash
   sed -i -E 's/newTag: ".*"/newTag: "0"/' k8s/base/kustomization.yaml
   ```
4. Commit and push:
   ```bash
   git add -A && git commit -m "chore: point pipelines at my fork" && git push
   ```

**Why it matters:** the pipelines' last stage commits the new image tag back to **your** repo, and Argo CD deploys from **your** repo.

**Optional – run the app locally first:** `docker compose up -d --build`, then open `http://localhost:3000` (shop) and `http://localhost:8000/playground`. Compose needs a `.env` with `POSTGRES_USER`, `POSTGRES_PASSWORD` and `POSTGRES_DB`.

---

## Step 2 – Create the Terraform state bucket

**Goal:** a place where Terraform stores what it created.

**Why:** Terraform remembers every resource it made in a *state file*. If that file lives on your laptop and you lose it, Terraform forgets your infrastructure. In S3 it is shared, versioned (you can recover old versions) and **locked** while someone runs `apply`, so two runs can't corrupt it.

**Do (laptop):**
```bash
export ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)

aws s3api create-bucket --bucket go-grpc-micro-tfstate-$ACCOUNT_ID \
  --region eu-north-1 --create-bucket-configuration LocationConstraint=eu-north-1

aws s3api put-bucket-versioning --bucket go-grpc-micro-tfstate-$ACCOUNT_ID \
  --versioning-configuration Status=Enabled
```
Bucket names are global across all AWS users, which is why the account ID is in the name.

**Check:** `aws s3 ls | grep tfstate`

---

## Step 3 – Create the Jenkins server

**Goal:** one EC2 instance running **Jenkins** (CI server) and **SonarQube** (code-quality server), plus the CLI tools the pipelines use (Docker, Trivy, Terraform, kubectl, Helm, AWS CLI).

**Why Terraform and not the console:** the server is described in code (`terraform/jenkins-server/`), so you can delete and recreate it identically.

**What gets created (12 resources):** a VPC with a public subnet, an internet gateway, a security group (8080 open for Jenkins and GitHub webhooks, 9000 only for your IP, **no SSH port**), an IAM role (admin + SSM), the EC2 instance (`t3.large`, Ubuntu 24.04, 30 GiB encrypted disk) and an Elastic IP (the address never changes).

**Do (laptop):**
```bash
cd terraform/jenkins-server

cp terraform.tfvars.example terraform.tfvars
# edit terraform.tfvars →  admin_cidrs = ["<YOUR_IP>/32"]

cp backend.hcl.example backend.hcl
# edit backend.hcl →  bucket = "go-grpc-micro-tfstate-<ACCOUNT_ID>"

terraform init -backend-config=backend.hcl   # downloads the AWS provider, connects to S3 state
terraform plan                               # shows "12 to add" – read it!
terraform apply                              # type "yes"
terraform output                             # jenkins_url, sonarqube_url, ssm_session
```

`terraform.tfvars` and `backend.hcl` are git-ignored on purpose: they contain your IP and account ID.

**What happens on first boot:** `setup.sh` runs automatically as *user data* and installs Java 21, Jenkins, Docker, SonarQube (as a container), AWS CLI, kubectl, eksctl, Terraform, Trivy and Helm. It takes 5–8 minutes.

**Watch it (laptop):**
```bash
aws ssm start-session --region eu-north-1 --target <JENKINS_INSTANCE_ID>   # from "terraform output ssm_session"
sudo tail -f /var/log/ec2-user-data.log      # wait for "INSTALLATION COMPLETE", then Ctrl+C
```
`aws ssm start-session` gives you a shell **without SSH**: no key pair and no open port 22. AWS checks your IAM permissions instead.

**Check:**
- `http://<JENKINS_IP>:8080` shows "Unlock Jenkins".
- `http://<JENKINS_IP>:9000` shows the SonarQube login.

**If it fails:**
- `TargetNotConnected` on start-session: the instance is still booting. Wait 1–2 minutes.
- SonarQube keeps restarting: check `sysctl vm.max_map_count` (must be 524288; `setup.sh` sets it).

---

## Step 4 – First-time Jenkins and SonarQube setup

### 4.1 Unlock Jenkins
```bash
# on the Jenkins server (SSM session)
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```
Paste it at `http://<JENKINS_IP>:8080`, choose **Install suggested plugins**, and create your admin user.

### 4.2 Install plugins
**Manage Jenkins → Plugins → Available plugins**, then install:
`SonarQube Scanner`, `Go`, `NodeJS`, `OWASP Dependency-Check`, `Docker Pipeline`, `Pipeline: AWS Steps`, `Pipeline: Stage View`.

### 4.3 SonarQube token and webhook
At `http://<JENKINS_IP>:9000`:
1. Log in as `admin` / `admin` and set a new password.
2. **Avatar → My Account → Security → Generate Tokens**: Name `jenkins`, Type **Global Analysis Token** → Generate → **copy it now** (it is shown only once).
3. **Administration → Configuration → Webhooks → Create**: Name `jenkins`, URL **`http://172.17.0.1:8080/sonarqube-webhook/`** (keep the trailing slash).

> **Why 172.17.0.1?** SonarQube runs in a Docker container on the same machine as Jenkins. From inside a container, `172.17.0.1` is the host machine. After a scan, SonarQube uses this webhook to tell Jenkins whether the *quality gate* passed.

### 4.4 Jenkins credentials (first two)
**Manage Jenkins → Credentials → System → Global credentials (unrestricted) → Add Credentials**

| Kind | ID | Secret |
|---|---|---|
| Secret text | `AWS_ACCOUNT_ID` | your account ID |
| Secret text | `sonar-token` | the SonarQube token |

> ⚠️ Always add credentials under **System → Global**. Credentials added under **your user** (click your name → Credentials) are invisible to pipelines, and the build fails with `ERROR: AWS_ACCOUNT_ID`.

You do **not** add AWS access keys: Jenkins gets AWS permissions from the EC2 instance's IAM role.

### 4.5 Connect Jenkins to SonarQube
**Manage Jenkins → System → SonarQube servers → Add**: Name **`sonar-server`**, URL **`http://localhost:9000`**, token `sonar-token`. Save.

> Use `localhost`, not the public IP. Port 9000 only allows your laptop's IP, so the server talking to its own public IP times out (`HTTP connect timed out`).

### 4.6 Tools (first two)
**Manage Jenkins → Tools**
- SonarQube Scanner installations → name **`sonar-scanner`**, Install automatically.
- Go installations → name **`go`**, version **1.26.x**.

---

## Step 5 – Create the EKS cluster from Jenkins

**Goal:** a Kubernetes cluster (EKS) in its own VPC, created by a Jenkins pipeline running Terraform.

**Why from Jenkins:** the infrastructure change runs in CI with a **plan → human approval → apply** flow, so you see exactly what will change before it happens.

**What gets created for `dev` (~46 resources):**
- a VPC `10.10.0.0/16` across 2 AZs, public + private subnets and 1 NAT gateway;
- EKS **1.36** with a **private-only API endpoint** (not reachable from the internet), secrets encrypted with a KMS key, and control-plane logs in CloudWatch;
- 2 worker nodes (`t3.large`, Amazon Linux 2023) in private subnets, IMDSv2 with hop limit 1 and encrypted disks;
- add-ons: VPC CNI, kube-proxy, CoreDNS, Pod Identity agent, EBS CSI driver (with a default `gp3` StorageClass);
- a **bastion** EC2 (no public IP, SSM only), the only machine allowed to run `kubectl`;
- AWS Secrets Manager secrets for the two databases, plus an IAM role for External Secrets Operator.

**Do (Jenkins UI):**
1. **New Item** → name `eks-cluster` → **Pipeline** → OK.
2. **Pipeline** section:
   - Definition: *Pipeline script from SCM* → SCM: **Git**
   - Repository URL: `https://github.com/<YOUR_GH_USER>/go-grpc-graphql-micro.git` (use https, not `git@…`)
   - Branch: `*/main`
   - Script Path: `terraform/eks-cluster/Jenkinsfile`
3. Save → **Build Now**. The first build only registers the parameters (and runs a `dev` plan).
4. **Build with Parameters** → `Environment = dev`, `Terraform_Action = apply` → Build.
5. Open **Console Output**. The pipeline pauses at **Approve**: read the plan (`Plan: 46 to add…`), then click **Proceed**.
6. Wait 15–20 minutes. At the end it prints the bastion command:
   ```text
   aws ssm start-session --region eu-north-1 --target i-0abc…
   ```

**What the Jenkinsfile does:**
```text
Init      terraform init -backend-config="bucket=go-grpc-micro-tfstate-${AWS_ACCOUNT_ID}" \
                         -backend-config="key=eks-cluster/${Environment}/terraform.tfstate"
Validate  terraform validate
Plan      terraform plan -var-file=envs/${Environment}.tfvars -out=tfplan   (-destroy for destroy)
Approve   input step – a human clicks Proceed
Apply     terraform apply tfplan      ← exactly the plan you approved
```

**Check (laptop):**
```bash
aws eks describe-cluster --region eu-north-1 --name go-microservices-dev \
  --query 'cluster.{status:status,version:version,public:resourcesVpcConfig.endpointPublicAccess}'
# { "status": "ACTIVE", "version": "1.36", "public": false }

aws secretsmanager list-secrets --region eu-north-1 --query 'SecretList[].Name'
# go-micro-shop/dev/account-db, go-micro-shop/dev/order-db
```

**If it fails:** `ERROR: AWS_ACCOUNT_ID` → see the warning in Step 4.4. "S3 bucket does not exist" → wrong account ID in the credential.

---

## Step 6 – Use the bastion to reach the cluster

**Why a bastion:** the Kubernetes API has **no public endpoint**. Only machines inside the VPC can reach it, and the bastion is the only one with admin rights. You reach the bastion with SSM, so there are no SSH keys and no open ports.

**Do (laptop):**
```bash
aws ssm start-session --region eu-north-1 --target <BASTION_INSTANCE_ID>
sudo su - ec2-user          # a normal home directory (SSM starts you as ssm-user in /)
```

**Check (bastion):**
```bash
kubectl get nodes -o wide          # 2 nodes Ready
kubectl get pods -A                # everything Running
kubectl get storageclass           # ebs-csi-default-sc (default)
```
`kubectl` is preinstalled and preconfigured (`/etc/eks/kubeconfig`), and `k` is an alias for it.

**Clone the repo on the bastion** (later steps use files from it):
```bash
git clone https://github.com/<YOUR_GH_USER>/go-grpc-graphql-micro.git
```

**If it fails:** `kubectl: command not found` right after connecting → run `exec bash -l`.

---

## Step 7 – Install the AWS Load Balancer Controller

**Goal:** when you create a Kubernetes **Ingress**, an AWS **Application Load Balancer** gets created automatically.

**Why:** EKS has no ingress controller by default. This controller watches Ingress objects and creates and configures ALBs. It needs AWS permissions, so we give it its own IAM role through **EKS Pod Identity**: only the controller's pods can use that role.

**Do (bastion, as ec2-user):**
```bash
mkdir -p ~/lbc && cd ~/lbc          # work in a folder you can write to

# Variables (re-run after every new session!)
cat > env.sh <<'EOF'
export CLUSTER_NAME=go-microservices-dev
export AWS_REGION=eu-north-1
export ACCOUNT_ID=$(aws sts get-caller-identity --query Account --output text)
export VPC_ID=$(aws eks describe-cluster --name $CLUSTER_NAME --region $AWS_REGION --query cluster.resourcesVpcConfig.vpcId --output text)
EOF
source env.sh
echo "$CLUSTER_NAME $AWS_REGION $ACCOUNT_ID $VPC_ID"     # all four must be filled in
```

**7.1 – IAM policy** (what the controller is allowed to do in AWS). Pin it to the same version as the controller:
```bash
curl -O https://raw.githubusercontent.com/kubernetes-sigs/aws-load-balancer-controller/v3.5.0/docs/install/iam_policy.json
aws iam create-policy --policy-name AWSLoadBalancerControllerIAMPolicy --policy-document file://iam_policy.json
```

**7.2 – IAM role** that EKS Pod Identity can assume:
```bash
cat > trust.json <<'EOF'
{
  "Version": "2012-10-17",
  "Statement": [{
    "Effect": "Allow",
    "Principal": { "Service": "pods.eks.amazonaws.com" },
    "Action": ["sts:AssumeRole", "sts:TagSession"]
  }]
}
EOF
aws iam create-role --role-name ${CLUSTER_NAME}-lb-controller-role --assume-role-policy-document file://trust.json
aws iam attach-role-policy --role-name ${CLUSTER_NAME}-lb-controller-role \
  --policy-arn arn:aws:iam::${ACCOUNT_ID}:policy/AWSLoadBalancerControllerIAMPolicy
```

**7.3 – Link the role to the controller's service account:**
```bash
aws eks create-pod-identity-association --region $AWS_REGION --cluster-name $CLUSTER_NAME \
  --namespace kube-system --service-account aws-load-balancer-controller \
  --role-arn arn:aws:iam::${ACCOUNT_ID}:role/${CLUSTER_NAME}-lb-controller-role
```
Any pod running as service account `aws-load-balancer-controller` in `kube-system` now gets temporary credentials for that role.

**7.4 – Install with Helm:**
```bash
helm repo add eks https://aws.github.io/eks-charts && helm repo update
helm install aws-load-balancer-controller eks/aws-load-balancer-controller \
  -n kube-system --version 3.5.0 \
  --set clusterName=$CLUSTER_NAME \
  --set serviceAccount.create=true \
  --set serviceAccount.name=aws-load-balancer-controller \
  --set region=$AWS_REGION \
  --set vpcId=$VPC_ID
```
`region` and `vpcId` are set explicitly because the nodes block pods from reading instance metadata (IMDS hop limit 1, a security setting). Without them the controller crashes with "failed to get VPC ID".

**Check:**
```bash
kubectl get deployment -n kube-system aws-load-balancer-controller        # 2/2 READY
kubectl get pod -n kube-system -l app.kubernetes.io/name=aws-load-balancer-controller \
  -o jsonpath='{.items[0].spec.containers[0].env[*].name}'; echo
# must include AWS_CONTAINER_CREDENTIALS_FULL_URI  → Pod Identity works
```

**Smoke test (creates and deletes a real ALB):**
```bash
kubectl create namespace alb-test
kubectl -n alb-test create deployment echo --image=nginxinc/nginx-unprivileged:stable-alpine --port=8080
kubectl -n alb-test expose deployment echo --port=80 --target-port=8080
kubectl -n alb-test create ingress echo --class=alb --rule="/*=echo:80" \
  --annotation alb.ingress.kubernetes.io/scheme=internet-facing \
  --annotation alb.ingress.kubernetes.io/target-type=ip
kubectl -n alb-test get ingress echo -w          # wait for ADDRESS, Ctrl+C
```
**Wait 2–3 minutes** (a new ALB's DNS name doesn't exist yet), then from your laptop run `curl -I http://<ADDRESS>/`, which should return `200 OK`. Clean up with `kubectl delete namespace alb-test`.

**If it fails:**

| Error | Fix |
|---|---|
| `curl: (23) … Permission denied` | you're in a read-only directory; `cd ~/lbc` |
| `AccessDenied … iam:CreatePolicy` | the bastion permission comes from Terraform; make sure Step 5 apply finished |
| `--role-name: expected one argument` | variables are empty; `source ~/lbc/env.sh` |
| `Could not resolve host` | the ALB is still provisioning; wait |
| `EntityAlreadyExists` (on a recreated cluster) | the policy already exists; continue with 7.2 |

---

## Step 8 – Install Argo CD

**Goal:** a GitOps tool inside the cluster that keeps the cluster in sync with your git repo.

**Why GitOps:** Jenkins never needs cluster credentials. It only changes a file in git; Argo CD notices the change and applies it. Rollback = `git revert`.

**Do (bastion):**
```bash
kubectl create namespace argocd
kubectl apply -n argocd --server-side --force-conflicts \
  -f https://raw.githubusercontent.com/argoproj/argo-cd/v3.5.3/manifests/install.yaml
kubectl get pods -n argocd -w          # all Running, Ctrl+C
```
`--server-side` is required for Argo CD 3.x; without it you get `metadata.annotations: Too long`.

**Expose the UI (internet-facing, only your IP):**
```bash
kubectl annotate svc argocd-server -n argocd service.beta.kubernetes.io/aws-load-balancer-scheme=internet-facing
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
kubectl patch svc argocd-server -n argocd -p '{"spec": {"loadBalancerSourceRanges": ["<YOUR_IP>/32"]}}'
kubectl get svc argocd-server -n argocd          # EXTERNAL-IP = hostname
```

**Log in:**
```bash
kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo
```
Open `https://<EXTERNAL-IP>` (accept the self-signed certificate warning) and log in as `admin`. **User Info → Update Password**, then:
```bash
kubectl -n argocd delete secret argocd-initial-admin-secret
```

---

## Step 9 – Create the SonarQube projects

**Goal:** two SonarQube projects whose keys match the Jenkinsfiles.

**Do:** in SonarQube, **Projects → Create Project → Manually** (or *Create a local project*):

| Display name / Project key | Main branch |
|---|---|
| `go-grpc-graphql-micro-backend` | `main` |
| `go-grpc-graphql-micro-frontend` | `main` |

For each: **Set Up → Locally → Use existing token** (your `jenkins` token) **→ Other (JS, TS, Go, Python, …) → Linux**. The sample `sonar-scanner` command it shows is **not** needed, because Jenkins runs the scan. If the wizard generates a new token, revoke it (**My Account → Security**); you don't need it.

---

## Step 10 – Create the ECR repositories

**Goal:** five private container registries, one per image.

**Why these settings:**
- `scanOnPush=true`: AWS also scans every image for CVEs.
- `IMMUTABLE`: a tag like `account:42` can never be overwritten, so you always know exactly what's running, and rollback is easy.

**Do (laptop, not the bastion, which has no ECR permissions):**
```bash
for repo in account catalog order graphql frontend; do
  aws ecr create-repository --region eu-north-1 --repository-name $repo \
    --image-scanning-configuration scanOnPush=true \
    --image-tag-mutability IMMUTABLE \
    --query 'repository.repositoryUri' --output text
done
```

**Check:**
```bash
aws ecr describe-repositories --region eu-north-1 \
  --query 'repositories[].[repositoryName,imageTagMutability,imageScanningConfiguration.scanOnPush]' --output table
```

---

## Step 11 – Remaining Jenkins setup (tools and credentials)

### 11.1 Tools
**Manage Jenkins → Tools**
- **NodeJS installations** → name **`nodejs`**, version 22 LTS (SonarQube needs Node.js to analyse TypeScript).
- **Dependency-Check installations** → name **`DP-Check`** → Install automatically → *Install from github.com*.

### 11.2 GitHub token
The pipelines push the new image tag back to your repo, so they need write access.

GitHub → **Settings → Developer settings → Personal access tokens → Fine-grained tokens → Generate new token**
- Repository access: **Only select repositories** → your fork. The permissions list only appears after this.
- **Permissions → + Add permissions → Contents** → change to **Read and write**.
- Expiration: 90 days.

(If you can't find it: classic token with only `public_repo`.)

### 11.3 Credentials (System → Global)

| Kind | ID | Value |
|---|---|---|
| Username with password | `github` | your GitHub username + the token |
| Secret text | `nvd-api-key` | your NVD API key |

You should now have **four**: `AWS_ACCOUNT_ID`, `sonar-token`, `github`, `nvd-api-key`.

### 11.4 Create the two jobs
Same as Step 5 (New Item → Pipeline → Pipeline script from SCM → https repo URL → `*/main`):

| Job name | Script Path |
|---|---|
| `backend` | `jenkins/Jenkinsfile-Backend` |
| `frontend` | `jenkins/Jenkinsfile-Frontend` |

---

## Step 12 – Install External Secrets Operator

**Goal:** database passwords flow from **AWS Secrets Manager** into Kubernetes automatically.

**Why:** nobody should ever type, see or store a database password. Terraform generated random passwords in Step 5 and wrote them **directly to Secrets Manager** (they are not even in the Terraform state). External Secrets Operator (ESO) runs in the cluster, reads them with a read-only IAM role, and creates Kubernetes Secrets the pods use.

**Do (bastion):**
```bash
helm repo add external-secrets https://charts.external-secrets.io && helm repo update
helm install external-secrets external-secrets/external-secrets \
  -n external-secrets --create-namespace --version 2.11.0 \
  --set installCRDs=true \
  --set serviceAccount.name=external-secrets
kubectl -n external-secrets get pods      # 3 Running
```
The service account name **must** be `external-secrets`, because Terraform created the Pod Identity association for exactly that name.

---

## Step 13 – Run the CI pipelines

**Goal:** build, test, scan and push all five images, then record their tags in git.

**Do:** Jenkins → **`backend` → Build Now**. Watch **Console Output** or the Stage View.

| Stage | What it does | Expect |
|---|---|---|
| Go Unit Tests | `go test` for the 4 services | quick |
| Sonarqube Analysis | sends code to SonarQube | `ANALYSIS SUCCESSFUL` |
| Quality Check | waits for the webhook | `Quality gate is 'OK'` within seconds |
| OWASP Dependency-Check | checks dependencies against the NVD | **first run: 10–20 min** (downloads ~400k CVEs); later runs are fast |
| Trivy File Scan | scans the source tree | report archived |
| Docker Image Build | builds 4 images | few minutes |
| TRIVY Image Scan | **fails the build on fixable CRITICAL CVEs** | usually passes |
| ECR Image Pushing | pushes `:<BUILD_NUMBER>` | images appear in ECR |
| Update Deployment file | bumps tags in `k8s/base/kustomization.yaml`, commits, pushes | commit `chore(deploy): update backend images to N` |

When `backend` succeeds, run **`frontend`**. Don't run both at the same time the very first time: they'd both download the OWASP database.

**Check:**
```bash
for r in account catalog order graphql frontend; do
  echo "$r: $(aws ecr list-images --region eu-north-1 --repository-name $r --query 'imageIds[].imageTag' --output text)"
done
git pull && sed -n '/^images:/,$p' k8s/base/kustomization.yaml     # tags are now your build numbers
```

**If it fails:**

| Error | Fix |
|---|---|
| `Failed to query server version … connect timed out` | SonarQube server URL must be `http://localhost:9000` (Step 4.5) |
| Quality Check waits 5 min then fails | webhook missing or wrong (Step 4.3) |
| `permission denied … /var/run/docker.sock` | Jenkins needs a restart to join the docker group: `http://<JENKINS_IP>:8080/safeRestart` |
| TRIVY Image Scan fails with a CRITICAL CVE | **the gate working.** Update the base image in the Dockerfile (e.g. newer `nginx:<ver>-alpine`), push, re-run |
| `Could not find credentials entry with ID …` | credential missing or in the wrong store (Step 4.4) |
| `git push` rejected | another pipeline pushed meanwhile; the pipeline already does `git pull --rebase`. Re-run it |

---

## Step 14 – Deploy with Argo CD

**Goal:** Argo CD deploys the whole app (databases, services, frontend, ingress) from `k8s/overlays/dev`.

**How the account ID stays out of git:** the manifests say just `image: account`, and `k8s/base/kustomization.yaml` holds only the tag. The script below creates the Argo CD Application and tells it the registry address (`<ACCOUNT_ID>.dkr.ecr…`), which it looks up live from AWS.

**Do (bastion):**
```bash
cd ~/go-grpc-graphql-micro && git pull
./k8s/scripts/create-argocd-app.sh dev
```

**Watch:**
```bash
kubectl -n argocd get application go-micro-shop-dev -w        # → Synced  Healthy
kubectl -n go-micro-shop get externalsecret                    # → SecretSynced
kubectl -n go-micro-shop get pods -w                           # 13 pods → Running
```
Or in the Argo CD UI → click the **go-micro-shop-dev** tile to see the tree.

**Normal things you'll see:** EBS volumes being created for the 3 databases; Elasticsearch taking 1–2 minutes; `catalog-service` restarting once or twice until Elasticsearch is up.

**Open the app:**
```bash
kubectl -n go-micro-shop get ingress go-micro-shop     # ADDRESS
```
Wait 2–3 minutes, then open `http://<ADDRESS>/` (shop) and `http://<ADDRESS>/playground` (GraphQL).

**Test the API from your laptop:**
```bash
curl -s -X POST -H 'Content-Type: application/json' \
  -d '{"query":"{ products(pagination:{skip:0,take:3}) { id name price } }"}' \
  http://<ADDRESS>/graphql
```

**If it fails:**

| Symptom | Command |
|---|---|
| ExternalSecret `SecretSyncedError` | `kubectl -n go-micro-shop describe externalsecret account-db-secret` |
| Pod `ImagePullBackOff` | `kubectl -n go-micro-shop describe pod <pod>`; are the images in ECR? Did you run the pipelines first? |
| PVC `Pending` | `kubectl -n go-micro-shop describe pvc <pvc>` |
| Ingress without ADDRESS | `kubectl -n kube-system logs deploy/aws-load-balancer-controller --tail=50` |
| Browser "DNS address could not be found" but curl works | browser cached the lookup: `chrome://net-internals/#dns` → Clear host cache, or use incognito |

---

## Step 15 – Prometheus and Grafana

**Goal:** see CPU, memory and pod health for the cluster and your app.

- **Prometheus** collects metrics (from nodes via node-exporter, from Kubernetes objects via kube-state-metrics, and from containers via cAdvisor).
- **Grafana** draws dashboards from them.

**Do (bastion):**
```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts && helm repo update
helm install prometheus prometheus-community/prometheus -n monitoring --create-namespace --version 29.35.0
helm install grafana oci://ghcr.io/grafana-community/helm-charts/grafana -n monitoring --version 13.2.7

# expose both, only to your IP (Prometheus has no login!)
kubectl -n monitoring patch svc prometheus-server -p '{"spec":{"type":"LoadBalancer","loadBalancerSourceRanges":["<YOUR_IP>/32"]}}'
kubectl -n monitoring patch svc grafana           -p '{"spec":{"type":"LoadBalancer","loadBalancerSourceRanges":["<YOUR_IP>/32"]}}'

kubectl -n monitoring get pods
kubectl -n monitoring get svc prometheus-server grafana                      # hostnames
kubectl -n monitoring get secret grafana -o jsonpath="{.data.admin-password}" | base64 -d; echo
```
Use the `grafana-community` chart; the old `grafana/grafana` chart is deprecated.

**Use them (after ~3 minutes, with `http://`):**
- Prometheus: `http://<prometheus-host>/` → **Status → Target health**, all UP.
- Grafana: `http://<grafana-host>/`, logging in as `admin` with the password above.
  1. **Connections → Data sources → Add → Prometheus**, URL `http://prometheus-server.monitoring.svc.cluster.local` → **Save & test**.
  2. **Dashboards → New → Import** → `6417` → Load → select Prometheus → Import. Repeat with `17375`.
  3. **Dashboards → New → Import** → upload `k8s/monitoring/grafana-dashboard.json` → Import. This is the app's own dashboard: requests, errors and latency per service, Postgres/Elasticsearch timings and Go runtime, read from each pod's `:9464/metrics` (pods carry `prometheus.io/scrape` annotations, so the chart's `kubernetes-pods` job finds them without extra config).

**If Grafana won't load in the browser but `curl` works:** you typed it without `http://` and the browser tried HTTPS, or the DNS lookup is cached. Also check that `https://checkip.amazonaws.com` in the **browser** shows `<YOUR_IP>` (a VPN changes it).

---

## Step 16 – SigNoz and OpenTelemetry

**Goal:** follow one request from the GraphQL gateway through the gRPC services down to the SQL query, and see the logs that request wrote.

- Every Go service uses the **OpenTelemetry SDK** (`telemetry/telemetry.go`): spans for GraphQL operations and resolvers, gRPC calls, Postgres queries and Elasticsearch requests; metrics; and `slog` logs carrying the trace ID.
- They send all three over **OTLP/gRPC** to the **SigNoz OTel Collector**, which stores them in ClickHouse. The address comes from the `otel-config` ConfigMap in `k8s/base/kustomization.yaml`.
- **Prometheus + Grafana stay.** The same metrics are also served on `:9464/metrics` and scraped by Prometheus (Step 15). SigNoz is for traces, logs and per-request debugging; Grafana keeps the cluster dashboards.
- If SigNoz is not installed, the services still run: OTLP exports just fail in the background.

**Do (bastion, from the repo root):**
```bash
helm repo add signoz https://charts.signoz.io && helm repo update
helm upgrade --install signoz signoz/signoz -n signoz --create-namespace \
  --version 0.145.0 -f k8s/monitoring/signoz-values.yaml

kubectl -n signoz get pods -w        # ClickHouse takes 3–5 minutes; Ctrl-C when all are Running

# the UI only, restricted to your IP (the chart's own Service would also publish internal ports)
MY_IP=$(curl -s https://checkip.amazonaws.com)
sed "s|<YOUR_IP>|${MY_IP}|" k8s/monitoring/signoz-ui-service.yaml | kubectl apply -f -
kubectl -n signoz get svc signoz-ui  # hostname
```
The release **must** be named `signoz` in namespace `signoz`: the services send to `signoz-otel-collector.signoz.svc.cluster.local:4317`.

Pods that started before SigNoz was up reconnect on their own. If a service shows no data after a few minutes, restart it: `kubectl -n go-micro-shop rollout restart deploy`.

**Use it (`http://<signoz-host>/`):**
1. Create the admin account on first visit.
2. Click around the shop (create an account, a product, an order) to generate traffic.
3. **Services:** `graphql-gateway`, `account-service`, `catalog-service`, `order-service` with request rate, error rate and p99.
4. **Traces:** open a `createOrder` trace. It shows the gateway → `order-service` → `account-service` / `catalog-service` → `INSERT` / Elasticsearch spans in one waterfall.
5. **Logs:** filter `service.name = order-service`; any error log links to its trace (**View trace**).
6. **Settings → General:** set retention (for example traces/logs 7 days, metrics 30 days) so ClickHouse's 20 Gi volume does not fill up.

**Check what a pod sends** (if something is missing):
```bash
kubectl -n go-micro-shop exec deploy/account-service -- wget -qO- localhost:9464/metrics | head   # metrics endpoint
kubectl -n go-micro-shop logs deploy/account-service | grep -i otlp                              # export errors
kubectl -n signoz logs deploy/signoz-otel-collector --tail=50
```

---

## Step 17 – Prove GitOps works

1. Change something visible in the frontend (e.g. a heading in `frontend/src/pages/Home.tsx`), commit and push.
2. Run the **`frontend`** job.
3. It pushes `frontend:<N>` and commits the new tag.
4. Within ~3 minutes Argo CD shows **OutOfSync → Syncing → Synced** and replaces the frontend pods. You never touched `kubectl`.

**Rollback:** `git revert <the chore(deploy) commit>` and push. Argo CD goes back to the previous image.

---

## Step 18 – Teardown

**Order matters.** Load balancers and EBS volumes created *by Kubernetes* must be deleted before Terraform deletes the VPC, otherwise the destroy hangs.

**Bastion:**
```bash
kubectl -n argocd delete application go-micro-shop-dev    # removes app, ALB, PVCs/EBS
helm uninstall grafana prometheus -n monitoring
helm uninstall signoz -n signoz
# the ClickHouse resource keeps a finalizer once its operator is gone; clear it so the namespace can go
kubectl -n signoz patch clickhouseinstallations.clickhouse.altinity.com/signoz-clickhouse \
  -p '{"metadata":{"finalizers":[]}}' --type=merge
kubectl -n signoz delete svc signoz-ui
kubectl delete namespace signoz                            # removes the SigNoz PVCs/EBS volumes
kubectl -n argocd delete svc argocd-server
kubectl get svc,ingress -A | grep -i loadbalancer          # must print nothing
```

**Jenkins:** `eks-cluster` → Build with Parameters → `dev` + `destroy` → Proceed.

**Laptop (optional, stops all costs):**
```bash
cd terraform/jenkins-server && terraform destroy
```

Left behind (cheap): ECR images, the state bucket, the `AWSLoadBalancerControllerIAMPolicy` policy and the `…-lb-controller-role` role.

---

## Appendix A – Glossary

| Term | Meaning |
|---|---|
| **Terraform state** | file where Terraform records what it created |
| **Backend / `backend.hcl`** | where state is stored (S3) / your local, git-ignored backend settings |
| **EKS** | AWS-managed Kubernetes |
| **Managed node group** | EC2 workers that AWS creates and updates for EKS |
| **Private endpoint** | the Kubernetes API is reachable only from inside the VPC |
| **Bastion** | a small server inside the VPC used as the admin entry point |
| **SSM Session Manager** | AWS feature for shell access without SSH |
| **IMDS / hop limit 1** | the EC2 metadata service; hop limit 1 stops containers from reaching it (and stealing the node's IAM role) |
| **Pod Identity** | gives a specific Kubernetes service account an IAM role |
| **Access entry** | maps an IAM user or role to Kubernetes permissions in EKS |
| **ALB / CLB / NLB** | Application / Classic / Network load balancers |
| **Ingress** | Kubernetes object describing HTTP routing; turned into an ALB here |
| **Kustomize base/overlay** | shared manifests / per-environment additions |
| **GitOps** | git is the source of truth; an agent (Argo CD) makes the cluster match it |
| **Quality gate** | SonarQube pass/fail rule for new code |
| **SCA / SAST** | dependency (software composition) analysis / static code analysis |
| **Immutable tag** | an image tag that can never be overwritten |
| **ESO** | External Secrets Operator: syncs secrets from AWS into Kubernetes |
| **Write-only argument** | Terraform sends the value to AWS but never stores it in state |

## Appendix B – Command cheat sheet

```bash
# Connect
aws ssm start-session --region eu-north-1 --target <INSTANCE_ID>

# Terraform (laptop)
terraform init -backend-config=backend.hcl [-backend-config="key=eks-cluster/dev/terraform.tfstate"]
terraform plan [-var-file=envs/dev.tfvars]
terraform apply | terraform destroy | terraform output

# Kubernetes (bastion)
kubectl get pods -n go-micro-shop -o wide
kubectl describe pod <pod> -n go-micro-shop
kubectl logs <pod> -n go-micro-shop [--previous]
kubectl get events -n go-micro-shop --sort-by=.lastTimestamp
kubectl get externalsecret,secretstore -n go-micro-shop
kubectl -n argocd get applications

# Helm
helm list -A
helm uninstall <release> -n <ns>

# AWS
aws ecr list-images --repository-name account
aws secretsmanager get-secret-value --secret-id go-micro-shop/dev/account-db --query SecretString --output text
aws elbv2 describe-load-balancers --query 'LoadBalancers[].[DNSName,State.Code]'
```

## Appendix C – Troubleshooting index

| Problem | Where |
|---|---|
| Credential not found in Jenkins | Step 4.4 |
| SonarQube timeout | Step 4.5 |
| Bastion command errors | Step 7 table |
| Pipeline failures | Step 13 table |
| Deployment failures | Step 14 table |
| Browser can't reach a new load balancer | Step 14, Step 15, Step 16 |
| Your IP changed | Update `terraform/jenkins-server/terraform.tfvars` + `terraform apply`; re-run the `loadBalancerSourceRanges` patches (Steps 8 and 15) with the new IP |
| `terraform destroy` hangs on the VPC | a load balancer or ENI still exists, see Step 18 |
