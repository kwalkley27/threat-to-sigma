package main

import (
	"context"
	"log"

	"github.com/kwalkley27/threat-to-sigma/config"
	"github.com/kwalkley27/threat-to-sigma/feeds"
	"github.com/kwalkley27/threat-to-sigma/genai"
)

func main() {
	ctx := context.Background()

	// Load default configs
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable or config value not set")
	}

	// Retrieve iocs from threat feed
	cidrList, err := feeds.Retrieve(cfg)
	if err != nil {
		log.Fatalf("Error retrieving IOCs: %v", err)
	}

	// Format sigma rules for retrieved iocs
	genai.FormatSigma(cfg, ctx, cidrList)
}
