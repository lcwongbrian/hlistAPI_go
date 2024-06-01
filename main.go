package main

import (
	"hlistAPI/dbConnector"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
	}

	mongoUri := os.Getenv("MONGO_URL")

	if len(mongoUri) == 0 {
		log.Fatal("Error loading environment variables")
	}

	dbConnector.Connect(os.Getenv("MONGO_URL"))
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), cors.Default())
	router.GET("/hlist/getSurfaceById/:id", dbConnector.GetSurfaceById)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
