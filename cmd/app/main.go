package main

import (
	"context"

	"github.com/kwalkley27/threat-to-sigma/genai"
	"github.com/kwalkley27/threat-to-sigma/feeds"
)

func main() {
	ctx := context.Background()

	// Retrieve iocs from threat feed
	cidrList := feeds.Retrieve()

	// Format sigma rules for retrieved iocs
	genai.FormatSigma(ctx, cidrList)
}