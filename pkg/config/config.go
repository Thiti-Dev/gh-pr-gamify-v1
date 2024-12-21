package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	GithubBearerToken string
}

// LoadConfig reads configuration from environment files and environment variables
func LoadConfig() (*Config, error) {
	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("failed to get working directory: %w", err)
	}

	// Setup Viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(workDir)

	// Read .env file if it exists
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Also read from .env.local if it exists (for local development)
	localEnvPath := filepath.Join(workDir, ".env.local")
	if _, err := os.Stat(localEnvPath); err == nil {
		viper.SetConfigName(".env.local")
		if err := viper.MergeInConfig(); err != nil {
			return nil, fmt.Errorf("failed to merge local config: %w", err)
		}
	}

	// Read from environment variables
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("GITHUB_BEARER_TOKEN", "")

	// Create config instance
	config := &Config{
		GithubBearerToken: viper.GetString("GITHUB_BEARER_TOKEN"),
	}

	// Validate required fields
	if config.GithubBearerToken == "" {
		return nil, fmt.Errorf("GITHUB_BEARER_TOKEN is required")
	}

	return config, nil
}
