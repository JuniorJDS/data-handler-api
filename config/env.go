package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSettings() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %v\n", err)
	}

	settings := make(map[string]string)

	settings["API_V1"] = "/api/v1"
	settings["PORT"] = "5000"

	settings["POSTGRES"] = os.Getenv("POSTGRES")
	return settings
}
