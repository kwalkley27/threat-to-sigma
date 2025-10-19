package genai

import (
	"fmt"
	"log"
	"context"
	"sync"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/kwalkley27/threat-to-sigma/config"
)

// Format the Gemini prompt with the given ioc
func getPrompt(ioc string) string {

	template := `You are a cyber threat intelligence analyst. 
	             Convert this indicator into a sigma rule: %v. 
				 Output only the sigma rule formatted exactly 
				 to sigma specifications.`

	return fmt.Sprintf(template, ioc)
}

// Generate a single Sigma translation
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

// Process definition used for asynchronous translation
func processIOC(ctx context.Context, model *genai.GenerativeModel, ioc string) {
	genSingleSigma(ctx, model, ioc)
	fmt.Println()
}

// Manages the overall Sigma formatting flow
func FormatSigma(cfg *config.Config, ctx context.Context, iocs []string) {
	
	//Load global configs
	//cfg := config.Load()
	
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

	//Generate and print sigma rule for each ioc asynchronously
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, cfg.MaxConcurrency)

	for _, ioc := range iocs {
		wg.Add(1)

		// Acquire semaphore slot
		semaphore <- struct{}{}

		go func(ioc string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release slot when done

			processIOC(ctx, model, ioc)
		}(ioc)
	}

	wg.Wait()


}