# Customized Go version

This directory is a Go-based starting point for Module 1. The application uses a configuration layer so the same code can target LocalStack or AWS.

## Run locally

Start LocalStack with Bedrock enabled and an Ollama-backed model, then run from this directory:

```powershell
go mod tidy
go run ./cmd/bedrock-workshop -prompt "What is generative AI?"
go run ./cmd/bedrock-workshop -stream -prompt "Explain RAG briefly"
```

The defaults are:

```text
BEDROCK_MODE=localstack
AWS_REGION=us-east-1
BEDROCK_ENDPOINT=http://localhost:4566
BEDROCK_MODEL_ID=ollama.llama3.2
```

For AWS, the Go code remains unchanged. Only the environment changes:

```powershell
$env:BEDROCK_MODE = "aws"
$env:BEDROCK_ENDPOINT = ""
$env:BEDROCK_MODEL_ID = "global.anthropic.claude-haiku-4-5-20251001-v1:0"
go run ./cmd/bedrock-workshop -stream
```

The Terraform environment is currently implemented only in `terraform/localstack`. The `terraform/aws` directory is reserved for the AWS implementation.