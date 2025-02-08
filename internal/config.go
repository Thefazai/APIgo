package internal

import (
	"log"
	"os"
	"strings"
)

func GetEnvVar(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic("ENVIRONMENT VARIABLE REQUIRED: " + key)
	}
	return value
}

func LoadConfig() {
	apiKey := GetEnvVar("APIGO_API_KEY") // No default

	// Basic validation
	if len(apiKey) < 32 {
		panic("Invalid API key format")
	}

	// Securely handle in memory
	defer func() {
		apiKey = strings.Repeat("x", len(apiKey))
	}()

	// Safe logging
	log.Printf("API Key initialized (%d characters)", len(apiKey))
}
