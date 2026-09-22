resource "aws_s3_bucket" "workshop" {
  bucket = "bedrock-workshop-local"
}

resource "aws_dynamodb_table" "sessions" {
  name         = "bedrock-workshop-sessions"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "session_id"

  attribute {
    name = "session_id"
    type = "S"
  }
}