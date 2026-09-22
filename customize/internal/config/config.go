package config

import (
	"fmt"
	"os"
)

type Config struct {
	Mode            string
	Region          string
	BedrockEndpoint string
	ModelID         string
}

func Load() Config {
	mode := getenv("BEDROCK_MODE", "localstack")
	region := getenv("AWS_REGION", "us-east-1")

	endpoint := os.Getenv("BEDROCK_ENDPOINT")
	modelID := os.Getenv("BEDROCK_MODEL_ID")
	if mode == "localstack" {
		if endpoint == "" {
			endpoint = "http://localhost:4566"
		}
		if modelID == "" {
			modelID = "ollama.llama3.2"
		}
	} else if modelID == "" {
		modelID = "global.anthropic.claude-haiku-4-5-20251001-v1:0"
	}

	return Config{
		Mode:            mode,
		Region:          region,
		BedrockEndpoint: endpoint,
		ModelID:         modelID,
	}
}

func (c Config) Validate() error {
	if c.Mode != "localstack" && c.Mode != "aws" {
		return fmt.Errorf("BEDROCK_MODE must be localstack or aws, got %q", c.Mode)
	}
	if c.Region == "" {
		return fmt.Errorf("AWS_REGION must not be empty")
	}
	if c.ModelID == "" {
		return fmt.Errorf("BEDROCK_MODEL_ID must not be empty")
	}
	return nil
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
