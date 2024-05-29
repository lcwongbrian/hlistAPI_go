package main

import (
	"hlistAPI/dbConnector"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	dbConnector.Connect()
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
