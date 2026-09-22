output "bedrock_mode" {
  value = "localstack"
}

output "bedrock_endpoint" {
  value = var.localstack_endpoint
}

output "aws_region" {
  value = var.aws_region
}

output "s3_bucket" {
  value = aws_s3_bucket.workshop.id
}

output "dynamodb_table" {
  value = aws_dynamodb_table.sessions.name
}