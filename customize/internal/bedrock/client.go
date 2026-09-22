package bedrock

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime/types"

	"github.com/aws-samples/amazon-bedrock-workshop/customize/internal/config"
)

type Client struct {
	config config.Config
	api    *bedrockruntime.Client
}

func New(ctx context.Context, cfg config.Config) (*Client, error) {
	awsCfg, err := awsconfig.LoadDefaultConfig(ctx, awsconfig.WithRegion(cfg.Region))
	if err != nil {
		return nil, fmt.Errorf("load AWS configuration: %w", err)
	}

	options := func(options *bedrockruntime.Options) {
		if cfg.BedrockEndpoint != "" {
			options.BaseEndpoint = aws.String(cfg.BedrockEndpoint)
		}
	}

	return &Client{config: cfg, api: bedrockruntime.NewFromConfig(awsCfg, options)}, nil
}

func (c *Client) Converse(ctx context.Context, prompt string) (string, error) {
	response, err := c.api.Converse(ctx, &bedrockruntime.ConverseInput{
		ModelId: aws.String(c.config.ModelID),
		Messages: []types.Message{{
			Role:    types.ConversationRoleUser,
			Content: []types.ContentBlock{&types.ContentBlockMemberText{Value: prompt}},
		}},
		InferenceConfig: &types.InferenceConfiguration{
			MaxTokens:   aws.Int32(512),
			Temperature: aws.Float32(0.7),
		},
	})
	if err != nil {
		return "", fmt.Errorf("converse: %w", err)
	}

	message, ok := response.Output.(*types.ConverseOutputMemberMessage)
	if !ok {
		return "", fmt.Errorf("converse returned an unsupported output type")
	}
	for _, block := range message.Value.Content {
		text, ok := block.(*types.ContentBlockMemberText)
		if ok {
			return text.Value, nil
		}
	}
	return "", fmt.Errorf("converse returned no text")
}

func (c *Client) ConverseStream(ctx context.Context, prompt string, write func(string) error) error {
	response, err := c.api.ConverseStream(ctx, &bedrockruntime.ConverseStreamInput{
		ModelId: aws.String(c.config.ModelID),
		Messages: []types.Message{{
			Role:    types.ConversationRoleUser,
			Content: []types.ContentBlock{&types.ContentBlockMemberText{Value: prompt}},
		}},
		InferenceConfig: &types.InferenceConfiguration{MaxTokens: aws.Int32(512)},
	})
	if err != nil {
		return fmt.Errorf("converse stream: %w", err)
	}

	stream := response.GetStream()
	for event := range stream.Events() {
		chunk, ok := event.(*types.ConverseStreamOutputMemberContentBlockDelta)
		if ok {
			text, textOK := chunk.Value.Delta.(*types.ContentBlockDeltaMemberText)
			if textOK {
				if err := write(text.Value); err != nil {
					return err
				}
			}
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("read stream: %w", err)
	}
	return stream.Close()
}
