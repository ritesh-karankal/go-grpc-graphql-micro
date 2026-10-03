output "instance_id" {
  description = "Jenkins EC2 instance ID"
  value       = aws_instance.jenkins.id
}

output "public_ip" {
  description = "Jenkins server Elastic IP"
  value       = aws_eip.jenkins.public_ip
}

output "jenkins_url" {
  description = "Jenkins UI"
  value       = "http://${aws_eip.jenkins.public_ip}:8080"
}

output "sonarqube_url" {
  description = "SonarQube UI"
  value       = "http://${aws_eip.jenkins.public_ip}:9000"
}

output "ssm_session" {
  description = "Command to open a shell on the Jenkins server"
  value       = "aws ssm start-session --region ${var.aws_region} --target ${aws_instance.jenkins.id}"
}
