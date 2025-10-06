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

func processIOC(ctx context.Context, model *genai.GenerativeModel, ioc string, wg *sync.WaitGroup) {
	defer wg.Done()
	genSingleSigma(ctx, model, ioc)
	fmt.Println()
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

	// Generate and print sigma rule for each ioc asynchronously
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, cfg.MaxConcurrency)

	for _,ioc := range iocs {
		wg.Add(1)
		go func(ioc string) {
			defer wg.Done()
			defer func() { <-semaphore }() // release slot when done

			processIOC(ctx, model, ioc, &wg)
		}(ioc) //avoid loop variable capture
	}

	wg.Wait()

}