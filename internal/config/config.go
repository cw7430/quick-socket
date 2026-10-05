package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Environment string

const (
	EnvDev   Environment = "dev"
	EnvStage Environment = "stage"
	EnvProd  Environment = "prod"
)

type Config struct {
	Env       Environment
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

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func Load() (*Config, error) {

	env := Environment(getEnv("APP_ENV", string(EnvDev)))

	switch env {
	case EnvDev:
		gin.SetMode(gin.DebugMode)

	case EnvStage, EnvProd:
		gin.SetMode(gin.ReleaseMode)

	default:
		return nil, fmt.Errorf("invalid APP_ENV: %s", env)
	}

	port := getEnv("PORT", "4000")

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
		Env:       env,
		Port:      port,
		APIURL:    apiURL,
		ClientURL: clientURL,
		APIKey:    apiKey,
	}, nil
}
