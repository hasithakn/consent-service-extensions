package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Config holds all application configuration
type Config struct {
	Port     string
	LogLevel string
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Load .env file if it exists
	loadEnvFile(".env")

	cfg := &Config{
		Port:     getEnv("PORT", "3001"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	return cfg
}

// getEnv gets an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// loadEnvFile loads environment variables from a .env file
func loadEnvFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		// .env file is optional, so just log and continue
		log.Printf("No %s file found, using environment variables or defaults", filename)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Only set if not already set in environment
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading %s: %v", filename, err)
	}
}
