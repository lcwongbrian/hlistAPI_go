package main

import (
	"hlistAPI/dbConnector"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	var errEnvRoot error
	dir, errPath := filepath.Abs(filepath.Dir(os.Args[0]))

	if errPath != nil {
		log.Fatal("Error getting root directory")
	}

	errEnvPath := godotenv.Load(filepath.Join(dir, ".env"))

	if errEnvPath != nil {
		errEnvRoot = godotenv.Load(filepath.Join("./", ".env"))
		if errEnvRoot != nil {
			log.Fatal("Error loading .env file")
		}
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
