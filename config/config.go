package config

import (
	"os"
)

type Config struct {
	GeminiAPIKey string //Google Gemini API key
	SpamhausFeedURL string //Spamhaus drop list URL
	FeedLimit int //Max entries to be retrieved from the threat feed on one run
	ModelName string //Gemini model to be used for genai inference
	MaxConcurrency int //max number of concurrent genai api calls
}

func Load() *Config {
	return &Config{
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		SpamhausFeedURL:   "https://www.spamhaus.org/drop/drop.txt",
		FeedLimit: 4,
		ModelName: "gemini-2.5-pro",
		MaxConcurrency: 2,
	}
}
