package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL   string
	Interval int
}

func Load() Config {
	err := godotenv.Load("./scripts/agent/.env")
	if err != nil {
		log.Println("arquivo .env não encontrado, usando ENV do sistema")
	}

	return Config{
		APIURL:   getEnv("API_URL", "http://localhost:8080/metrics"),
		Interval: getEnvInt("INTERVAL", 10),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
