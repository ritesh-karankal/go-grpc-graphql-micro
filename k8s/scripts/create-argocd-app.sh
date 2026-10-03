#!/bin/bash
# Create (or update) the Argo CD Application for one environment.
# Run on the bastion. The ECR registry contains the AWS account ID, so it is
# looked up here and set on the Application instead of being stored in git.
#
# Usage: ./create-argocd-app.sh dev|prod

set -euo pipefail

ENVIRONMENT="${1:?usage: $0 dev|prod}"
REPO_URL="https://github.com/ritesh-karankal/go-grpc-graphql-micro.git"
AWS_REGION="eu-north-1"

ACCOUNT_ID="$(aws sts get-caller-identity --query Account --output text)"
REGISTRY="${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"

kubectl apply -f - <<YAML
apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: go-micro-shop-${ENVIRONMENT}
  namespace: argocd
spec:
  project: default
  source:
    repoURL: ${REPO_URL}
    targetRevision: main
    path: k8s/overlays/${ENVIRONMENT}
    kustomize:
      # Registry only: tags come from k8s/base/kustomization.yaml (bumped by Jenkins)
      images:
        - account=${REGISTRY}/account
        - catalog=${REGISTRY}/catalog
        - order=${REGISTRY}/order
        - graphql=${REGISTRY}/graphql
        - frontend=${REGISTRY}/frontend
  destination:
    server: https://kubernetes.default.svc
    namespace: go-micro-shop
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
      - CreateNamespace=true
YAML

echo "Application go-micro-shop-${ENVIRONMENT} applied. Watch it with:"
echo "  kubectl -n argocd get application go-micro-shop-${ENVIRONMENT} -w"
