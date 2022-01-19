package config

import (
	"github.com/joho/godotenv"
	"log"
)

func GetAppConfig() (appConfig map[string]string) {
	appConfig, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
		return nil
	}

	return
}
