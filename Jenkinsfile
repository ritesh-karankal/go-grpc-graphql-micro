pipeline {
    agent any

    tools {
        go 'go'
    }

    environment {
        AWS_REGION         = 'eu-north-1'
        AWS_ACCOUNT_ID     = credentials('AWS_ACCOUNT_ID')
        SONAR_SERVER       = 'sonar-server'
        SCANNER_HOME       = tool 'sonar-scanner'
        ECR_REGISTRY       = "${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"
        IMAGE_TAG          = "${BUILD_NUMBER}"
    }

    stages {
        stage('1. Clean Workspace') {
            steps {
                cleanWs()
            }
        }

        stage('2. Checkout Source Code') {
            steps {
                checkout scm
            }
        }

        stage('3. Go Unit Tests') {
            steps {
                sh '''
                    echo "=========================================="
                    echo "Running Go Unit Tests..."
                    echo "=========================================="
                    go test -v ./...
                '''
            }
        }

        stage('4. SonarQube Static Code Analysis') {
            steps {
                withSonarQubeEnv("${SONAR_SERVER}") {
                    sh '''
                        echo "=========================================="
                        echo "Running SonarQube Static Analysis..."
                        echo "=========================================="
                        $SCANNER_HOME/bin/sonar-scanner \
                          -Dsonar.projectName=go-grpc-graphql-micro \
                          -Dsonar.projectKey=go-grpc-graphql-micro \
                          -Dsonar.sources=. \
                          -Dsonar.exclusions=**/*_test.go,**/vendor/**,**/pb/**,node_modules/**
                    '''
                }
            }
        }

        stage('5. Quality Gate Check') {
            steps {
                timeout(time: 5, unit: 'MINUTES') {
                    waitForQualityGate abortPipeline: true
                }
            }
        }

        stage('6. Filesystem Vulnerability Scan (Trivy)') {
            steps {
                sh '''
                    echo "=========================================="
                    echo "Scanning Filesystem for Vulnerabilities..."
                    echo "=========================================="
                    trivy fs --severity HIGH,CRITICAL --exit-code 0 .
                '''
            }
        }

        stage('7. Build Microservice Docker Images') {
            steps {
                script {
                    def services = ['account', 'catalog', 'order', 'graphql', 'frontend']
                    for (service in services) {
                        def dockerfilePath = (service == 'frontend') ? 'frontend/Dockerfile' : "${service}/app.dockerfile"
                        def buildContext   = (service == 'frontend') ? 'frontend' : '.'
                        
                        echo "=========================================="
                        echo "Building Docker image: ${service}"
                        echo "=========================================="
                        sh "docker build -t ${ECR_REGISTRY}/${service}:${IMAGE_TAG} -t ${ECR_REGISTRY}/${service}:latest -f ${dockerfilePath} ${buildContext}"
                    }
                }
            }
        }

        stage('8. Trivy Container Security Scan') {
            steps {
                script {
                    def services = ['account', 'catalog', 'order', 'graphql', 'frontend']
                    for (service in services) {
                        echo "=========================================="
                        echo "Scanning Container Image: ${service}"
                        echo "=========================================="
                        sh "trivy image --severity HIGH,CRITICAL --exit-code 0 ${ECR_REGISTRY}/${service}:${IMAGE_TAG}"
                    }
                }
            }
        }

        stage('9. ECR Login & Push Images') {
            steps {
                script {
                    echo "=========================================="
                    echo "Authenticating with AWS ECR in ${AWS_REGION}..."
                    echo "=========================================="
                    sh "aws ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}"
                    
                    def services = ['account', 'catalog', 'order', 'graphql', 'frontend']
                    for (service in services) {
                        echo "=========================================="
                        echo "Ensuring ECR Repo & Pushing Image: ${service}"
                        echo "=========================================="
                        sh "aws ecr describe-repositories --repository-names ${service} --region ${AWS_REGION} || aws ecr create-repository --repository-name ${service} --region ${AWS_REGION}"
                        sh "docker push ${ECR_REGISTRY}/${service}:${IMAGE_TAG}"
                        sh "docker push ${ECR_REGISTRY}/${service}:latest"
                    }
                }
            }
        }
    }

    post {
        always {
            echo "Pipeline execution completed."
            cleanWs()
        }
        success {
            echo "🎉 DevSecOps Pipeline Completed Successfully! All microservices pushed to AWS ECR."
        }
        failure {
            echo "❌ Pipeline Failed! Check SonarQube or Trivy scan logs."
        }
    }
}
