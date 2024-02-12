package main

import (
	"log"
	"singnalzero-assesment/repository"
	"singnalzero-assesment/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	server := gin.Default()
	err := repository.ConnectToDatabase()

	if err != nil {
		log.Fatal("Error while connecting db", err)
	}

	routes.RegisterRoutes(server)

	server.Run(":8080")

}
