package dbConnector

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	mongoClient = client
	if err != nil {
		log.Fatal("Could not connect to MongoDB")
	}
}

func GetSurfaceById(c *gin.Context) {
	idParam := c.Param("id")
	var result bson.M
	id, errAtoi := strconv.Atoi(idParam)
	if errAtoi != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": errAtoi.Error()})
		return
	}

	errQuery := mongoClient.Database("spatial").Collection("hlist").FindOne(context.TODO(), bson.D{{Key: "surface_id", Value: id}}).Decode(&result)
	if errQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errQuery.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}
