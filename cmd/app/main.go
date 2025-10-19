package main

import (
	"context"

	"github.com/kwalkley27/threat-to-sigma/genai"
	"github.com/kwalkley27/threat-to-sigma/feeds"
	"github.com/kwalkley27/threat-to-sigma/config"
)

func main() {
	ctx := context.Background()

	// Load default configs
	cfg := config.Load()

	// Retrieve iocs from threat feed
	cidrList := feeds.Retrieve(cfg)

	// Format sigma rules for retrieved iocs
	genai.FormatSigma(cfg, ctx, cidrList)
}