package main

import (
	"log"
	"os"

	"github.com/jaddek/ecb/converter"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Overload(".env", ".env.local")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	host := os.Getenv("HOST_ECB_EU_RATES")

	log.Println(string(converter.Convert(host)))
}
