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
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	MongoURI      = "mongodb://admin:password@localhost:27017/test?authSource=admin"
	MongoDatabase = "demo"
)

func mongoConnect(ctx context.Context) (client *mongo.Client, err error) {
	client, err = mongo.Connect(options.Client().ApplyURI(MongoURI))
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	return
}

func mongoDisconnect(ctx context.Context, client *mongo.Client) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatal(err)
	}
}

func redisConnect() *redis.Client {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	status := redisClient.Ping(ctx)
	log.Println("Redis status:", status)
	return redisClient
}

func main() {
	ctx := context.Background()
	client, err := mongoConnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoDisconnect(ctx, client)

	redisClient := redisConnect()
	defer redisClient.Close()

	recipesH := handlers.NewRecipesHandler(
		client.Database(MongoDatabase).Collection("recipes"),
		redisClient,
	)
	authH := handlers.AuthHandler{}

	router := gin.Default()

	// auth
	router.POST("/signin", authH.SignInHandler)
	router.POST("/refresh", authH.RefreshHandler)

	authorized := router.Group("/")
	authorized.Use(authH.AuthMiddleware())

	// define the routes
	authorized.POST("/recipes", recipesH.CreateRecipeHandler)
	authorized.GET("/recipes", recipesH.ListRecipesHandler)
	authorized.PUT("/recipes/:id", recipesH.UpdateRecipeHandler)
	authorized.DELETE("/recipes/:id", recipesH.DeleteRecipeHandler)

	router.GET("/recipes/search", recipesH.SearchRecipesHandler)

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
