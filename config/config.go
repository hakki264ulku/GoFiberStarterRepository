package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config Gets the env value from .env file
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
		panic(err)
	}
	// Return the value of the variable
	return os.Getenv(key)
}
