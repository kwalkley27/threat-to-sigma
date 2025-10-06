package config

import (
	"os"
)

type Config struct {
	GeminiAPIKey string
	SpamhausFeedURL string
	FeedLimit int
	ModelName string
	MaxConcurrency int
}

func Load() *Config {
	return &Config{
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		SpamhausFeedURL:   "https://www.spamhaus.org/drop/drop.txt",
		FeedLimit: 2,
		ModelName: "gemini-2.5-pro",
		MaxConcurrency: 5,
	}
}
