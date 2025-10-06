package genai

import (
	"fmt"
	"os"
	"log"
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const ModelName = "gemini-2.5-pro"

func FormatSigma(ctx context.Context, iocs []string) {
		apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}

	// Create a new client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Choose a model (gemini-1.5-flash is fast & cheap, gemini-1.5-pro is more powerful)
	model := client.GenerativeModel(ModelName)
	
	prompt := fmt.Sprintf("You are a cyber threat intelligence analyst. Convert this indicator into a sigma rule: %v. Output only the sigma rule formatted exactly to sigma specifications", iocs[0])

	// Send a simple text prompt
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err!=nil {
		log.Fatalf("Request failed: %v", err)
	}

	// Print response
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			fmt.Println(part)
		}
	}
}