package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/aws-samples/amazon-bedrock-workshop/customize/internal/bedrock"
	"github.com/aws-samples/amazon-bedrock-workshop/customize/internal/config"
)

func main() {
	stream := flag.Bool("stream", false, "stream the response as it is generated")
	prompt := flag.String("prompt", "Explain Amazon Bedrock in one paragraph.", "prompt to send")
	flag.Parse()

	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		fatal(err)
	}

	ctx := context.Background()
	client, err := bedrock.New(ctx, cfg)
	if err != nil {
		fatal(err)
	}

	fmt.Printf("mode=%s region=%s model=%s endpoint=%s\n", cfg.Mode, cfg.Region, cfg.ModelID, cfg.BedrockEndpoint)
	if *stream {
		err = client.ConverseStream(ctx, *prompt, func(text string) error {
			_, writeErr := fmt.Print(text)
			return writeErr
		})
		fmt.Println()
	} else {
		var answer string
		answer, err = client.Converse(ctx, *prompt)
		if err == nil {
			fmt.Println(answer)
		}
	}
	if err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
