package main

import (
	"fmt"
	"os"
	"log"
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const ModelName = "gemini-2.5-flash"

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}

	ctx := context.Background()

	// Create a new client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Choose a model (gemini-1.5-flash is fast & cheap, gemini-1.5-pro is more powerful)
	model := client.GenerativeModel(ModelName)

	// Send a simple text prompt
	resp, err := model.GenerateContent(ctx, genai.Text("Write a haiku about Go programming."))
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