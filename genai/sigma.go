package genai

import (
	"fmt"
	"log"
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/kwalkley27/threat-to-sigma/config"
)

func getPrompt(s string) string {

	template := `You are a cyber threat intelligence analyst. 
	             Convert this indicator into a sigma rule: %v. 
				 Output only the sigma rule formatted exactly 
				 to sigma specifications.`

	return fmt.Sprintf(template, s)
}

func genSingleSigma(ctx context.Context, model *genai.GenerativeModel, ioc string) {
	
	// Generate prompt from ioc
	prompt := getPrompt(ioc)

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

func FormatSigma(ctx context.Context, iocs []string) {
	
	//Load global configs
	cfg := config.Load()
	
	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable not set")
	}

	// Create a new client
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiAPIKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Set Model
	model := client.GenerativeModel(cfg.ModelName)
	
	// Generate and print sigma rule for each ioc
	for _,ioc := range iocs {
		genSingleSigma(ctx, model, ioc)
		fmt.Println()
	}

}