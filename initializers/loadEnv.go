package initializers

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		message := fmt.Sprintf("Error loading .env file! Error: %s", err.Error())
		log.Fatal(message)
	}
}
