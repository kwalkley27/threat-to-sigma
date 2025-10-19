package genai

import (
	"context"
	"fmt"
	"log"
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
func genSingleSigma(ctx context.Context, model *genai.GenerativeModel, ioc string) (string, error) {
	prompt := getPrompt(ioc)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	var result string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			result += fmt.Sprintf("%v", part)
		}
	}
	return result, nil
}

// Process definition used for asynchronous translation
func processIOC(ctx context.Context, model *genai.GenerativeModel, ioc string, results chan<- string, errs chan<- error) {
	sigma, err := genSingleSigma(ctx, model, ioc)
	if err != nil {
		errs <- fmt.Errorf("error processing IOC %s: %w", ioc, err)
		return
	}
	results <- sigma
}

// Manages the overall Sigma formatting flow
func FormatSigma(cfg *config.Config, ctx context.Context, iocs []string) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(cfg.GeminiAPIKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel(cfg.ModelName)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, cfg.MaxConcurrency)
	results := make(chan string, len(iocs))
	errs := make(chan error, len(iocs))

	for _, ioc := range iocs {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(ioc string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			processIOC(ctx, model, ioc, results, errs)
		}(ioc)
	}

	wg.Wait()
	close(results)
	close(errs)

	for err := range errs {
		log.Printf("Error: %v", err)
	}

	for result := range results {
		fmt.Println(result)
		fmt.Println()
	}
}
