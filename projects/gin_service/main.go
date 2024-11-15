// Recipes API
//
// This is a sample recipes API.
//
//	Schemes: http
//	Host: localhost:8080
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	"context"
	"gin_service/handlers"
	"gin_service/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	MongoURI      = "mongodb://admin:password@localhost:27017/test?authSource=admin"
	MongoDatabase = "demo"
)

var recipes []models.Recipe
var client *mongo.Client

func mongoConnect() (err error) {
	client, err = mongo.Connect(options.Client().ApplyURI(MongoURI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func mongoDisconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := mongoConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDisconnect()

	handlers := handlers.NewRecipesHandler(client.Database(MongoDatabase).Collection("recipes"))

	router := gin.Default()

	// define the routes
	router.GET("/recipes", handlers.ListRecipesHandler)
	router.POST("/recipes", handlers.CreateRecipeHandler)
	router.PUT("/recipes/:id", handlers.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", handlers.DeleteRecipeHandler)
	router.GET("/recipes/search", handlers.SearchRecipesHandler)

	// run the server
	err = router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func IndexHandler(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello " + name,
	})
}
