package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
// 	fmt.Println("Received body: ", request.Body)

// 	converted := tool.Convert()

// 	return events.APIGatewayProxyResponse{Body: string(converted), StatusCode: 200}, nil
// }

func init() {
	err := godotenv.Overload(".env", ".env.local")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	host := os.Getenv("HOST_ECB_EU_RATES")

	log.Println(host)
}
