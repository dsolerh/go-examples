package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_service/models"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

const (
	MongoURI      = "mongodb://admin:password@localhost:27017/test?authSource=admin"
	MongoDatabase = "demo"
)

func main() {
	client, err := mongo.Connect(options.Client().ApplyURI(MongoURI))
	if err != nil {
		panic(err)
	}

	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connect successful")

	log.Println("getting all recipes")
	getAllRecipes(client)
}

func insertManyRecipes(client *mongo.Client) {
	recipes := make([]models.Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal(file, &recipes)

	listOfRecipes := make([]any, 0)
	for _, recipe := range recipes {
		listOfRecipes = append(listOfRecipes, recipe)
	}
	collection := client.Database(MongoDatabase).Collection("recipes")
	ctx := context.Background()
	insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Awk: ", insertManyResult.Acknowledged)
	log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))
}

func getAllRecipes(client *mongo.Client) {
	collection := client.Database(MongoDatabase).Collection("recipes")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	for cursor.Next(context.Background()) {
		var recipe models.Recipe
		err = cursor.Decode(&recipe)
		if err != nil {
			panic(err)
		}
		log.Println("Recipe: ", recipe.Name)
	}
}
