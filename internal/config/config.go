package config

import (
	"errors"
	"os"
	"time"
)

// Config holds application configuration
type Config struct {
	// Unsplash API access key
	UnsplashAccessKey string

	// API request timeout
	RequestTimeout time.Duration
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	accessKey := os.Getenv("UNSPLASH_ACCESS_KEY")
	if accessKey == "" {
		return nil, errors.New("UNSPLASH_ACCESS_KEY environment variable must be set")
	}

	return &Config{UnsplashAccessKey: accessKey, RequestTimeout: 10 * time.Second}, nil
}
