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

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
