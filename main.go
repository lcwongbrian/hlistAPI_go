package main

import (
	"context"
	"hlistAPI/dbConnector"
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file")
	}

	router := gin.New()
	router.Use(gin.Recovery(), cors.Default())

	dbHandler, err := dbConnector.NewDBHandler()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer dbHandler.Close()

	ctx := context.Background()

	err = dbHandler.Initialize(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}

	router.GET("/hlist/getSurfaceById/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid ID"})
			return
		}

		result, err := dbHandler.FetchSurfaceBinaryByID(ctx, int32(id))
		if err != nil {
			c.JSON(500, gin.H{"error": "Database error: " + err.Error()})
			return
		}

		c.JSON(200, result)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = router.Run(":" + port)
	if err != nil {
		log.Panicf("error: %s", err)
	}
}
