package config

import (
	"os"
)

type Config struct {
	GeminiAPIKey string
	SpamhausFeedURL string
	FeedLimit int
	ModelName string
}

func Load() *Config {
	return &Config{
		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),
		SpamhausFeedURL:   "https://www.spamhaus.org/drop/drop.txt",
		FeedLimit: 5,
		ModelName: "gemini-2.5-pro",
	}
}
