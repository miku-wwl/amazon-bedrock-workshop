# LocalStack environment

This Terraform configuration creates the local S3 and DynamoDB resources used by the Go customization project.

Start LocalStack on port `4566`, then run:

```powershell
terraform init
terraform apply
```

The application configuration is selected independently through environment variables:

```powershell
$env:BEDROCK_MODE = "localstack"
$env:BEDROCK_ENDPOINT = "http://localhost:4566"
$env:BEDROCK_MODEL_ID = "ollama.llama3.2"
go run ./cmd/bedrock-workshop -stream -prompt "Say hello"
```

`terraform/aws` is intentionally empty for now. It will later contain the AWS implementation without changing the Go configuration contract.