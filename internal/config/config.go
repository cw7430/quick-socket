package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port      string
	APIURL    string
	ClientURL string
	APIKey    string
}

func requiredEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is required", key)
	}

	return value, nil
}

func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	apiURL, err := requiredEnv("API_URL")
	if err != nil {
		return nil, err
	}

	clientURL, err := requiredEnv("CLIENT_URL")
	if err != nil {
		return nil, err
	}

	apiKey, err := requiredEnv("API_KEY")
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:      port,
		APIURL:    apiURL,
		ClientURL: clientURL,
		APIKey:    apiKey,
	}, nil
}
