package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	GeminiAPIKey    string `mapstructure:"gemini_api_key"`
	SpamhausFeedURL string `mapstructure:"spamhaus_feed_url"`
	FeedLimit       int    `mapstructure:"feed_limit"`
	ModelName       string `mapstructure:"model_name"`
	MaxConcurrency  int    `mapstructure:"max_concurrency"`
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// sensible defaults so the program can run without a config file
	viper.SetDefault("spamhaus_feed_url", "https://www.spamhaus.org/drop/drop.txt")
	viper.SetDefault("feed_limit", 10)
	viper.SetDefault("model_name", "gemini-1.5-pro")
	viper.SetDefault("max_concurrency", 5)

	// Try to read config file, but do not fail if it's missing — allow
	// environment variables and defaults to drive configuration.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// No config file found; continue with env vars and defaults
		} else {
			return nil, err
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
