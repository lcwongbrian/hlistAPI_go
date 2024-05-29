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
		log.Fatal("Error loading .env file")
	}
	dbConnector.Connect(os.Getenv("MONGO_URL"))
}

func main() {
	router := gin.New()
	router.Use(gin.Recovery(), cors.Default())
	router.GET("/hlist/getSurfaceById/:id", dbConnector.GetSurfaceById)
	err := router.Run("localhost:8080")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
