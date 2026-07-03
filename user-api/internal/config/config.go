package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort      string
	ShutdownTimeout time.Duration
}

func LoadConfig() AppConfig {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}

	timeoutStr := os.Getenv("SHUTDOWN_TIMEOUT")
	timeoutInt, err := strconv.Atoi(timeoutStr)
	if err != nil {
		timeoutInt = 5
	}

	return AppConfig{
		ServerPort:      port,
		ShutdownTimeout: time.Duration(timeoutInt) * time.Second,
	}
}
