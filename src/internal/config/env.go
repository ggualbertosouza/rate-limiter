package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Env        string // GO_ENV
	ServerPort string // SERVER_PORT
	ServerHost string // SERVER_HOST
}

func LoadEnvConfig() *EnvConfig {
	env := getEnv("GO_ENV", "development")
	if env == "development" {
		err := godotenv.Load(".env.development")
		if err != nil {
			log.Printf("failed to load .env.development")
		}
	}

	return &EnvConfig{
		Env:        env,
		ServerHost: getEnv("SERVER_HOST", "localhost"),
		ServerPort: ":" + strconv.Itoa(getEnvInt("SERVER_PORT", 8080)),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		return defaultValue
	}

	return defaultValue
}
