package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// FileDirectory holds the directory path for file operations
var (
	FileDirectory string
	PORT          string
)

// LoadConfig loads configuration from environment variables
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the directory from environment variable
	FileDirectory = os.Getenv("FILE_DIRECTORY")
	if FileDirectory == "" {
		log.Fatal("FILE_DIRECTORY not set in the environment")
	}

	// Get the port from the environment variable, default to 8080
	PORT = os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("PORT not set in the environment")
	}
}
