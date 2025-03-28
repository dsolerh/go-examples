package handlers

import (
	"encoding/json"
	"gin_service/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RecipesHandler struct {
	collection  *mongo.Collection
	redisClient *redis.Client
}

func NewRecipesHandler(collection *mongo.Collection, redisClient *redis.Client) RecipesHandler {
	return RecipesHandler{
		collection:  collection,
		redisClient: redisClient,
	}
}

// swagger:operation POST /recipes recipes createRecipe
// Returns the newly created recipe
// ---
// parameters:
//   - name: name
//     in: body
//     description: name of the recipe
//     required: true
//     type: string
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//		description: Successful operation
func (h RecipesHandler) CreateRecipeHandler(c *gin.Context) {
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = bson.NewObjectID()
	recipe.PublishedAt = time.Now()

	result, err := h.collection.InsertOne(c.Request.Context(), recipe)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Println("Insert result", result)

	if err := h.redisClient.Del(c.Request.Context(), "recipes").Err(); err != nil {
		log.Println("Error deleting the redis cache:", err.Error())
	}

	c.JSON(http.StatusOK, recipe)
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
//   - application/json
//
// responses:
//
//	'200':
//		description: Successful operation
func (h RecipesHandler) ListRecipesHandler(c *gin.Context) {
	val, err := h.redisClient.Get(c.Request.Context(), "recipes").Result()
	if err == redis.Nil {
		log.Println("Request to MongoDB")
		cur, err := h.collection.Find(c.Request.Context(), bson.D{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cur.Close(c.Request.Context())

		recipes := make([]models.Recipe, 0)
		for cur.Next(c.Request.Context()) {
			var recipe models.Recipe
			err = cur.Decode(&recipe)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			recipes = append(recipes, recipe)
		}

		// save to redis
		if recipesData, err := json.Marshal(&recipes); err == nil {
			err = h.redisClient.Set(c.Request.Context(), "recipes", recipesData, 0).Err()
			if err != nil {
				log.Println("redis could not set", err.Error())
			}
		} else {
			log.Println("could not serialize the recipes", err.Error())
		}

		c.JSON(http.StatusOK, recipes)
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		log.Println("Request to redis")

		recipes := make([]models.Recipe, 0)
		err := json.Unmarshal([]byte(val), &recipes)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, recipes)
	}
}

// swagger:operation PUT /recipes/{id} recipes updateRecipe
// Update an existing recipe
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//		description: Successful operation
//	'400':
//		description: Invalid input
//	'404':
//		description: Invalid recipe ID
func (h RecipesHandler) UpdateRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	var recipe models.Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recipeID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.collection.UpdateByID(c.Request.Context(), recipeID, bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: recipe.Name},
		{Key: "instructions", Value: recipe.Instructions},
		{Key: "ingredients", Value: recipe.Ingredients},
		{Key: "tags", Value: recipe.Tags},
	}}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	if err := h.redisClient.Del(c.Request.Context(), "recipes").Err(); err != nil {
		log.Println("Error deleting the redis cache:", err.Error())
	}

	c.JSON(http.StatusOK, recipe)
}

// swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// Deletes the specified recipe
// ---
// parameters:
//   - name: id
//     in: path
//     description: ID of the recipe
//     required: true
//     type: string
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//		description: Successful operation
//	'404':
//		description: Invalid recipe ID
func (h RecipesHandler) DeleteRecipeHandler(c *gin.Context) {
	id := c.Param("id")
	recipeID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.collection.DeleteOne(c.Request.Context(), bson.D{{Key: "_id", Value: recipeID}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
		return
	}

	if err := h.redisClient.Del(c.Request.Context(), "recipes").Err(); err != nil {
		log.Println("Error deleting the redis cache:", err.Error())
	}

	// recipes = append(recipes[:index], recipes[index+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
}

// swagger:operation GET /recipes/search recipes searchRecipes
// Searches the recipes by tag
// ---
// parameters:
//   - name: tag
//     in: query
//     description: the search tag
//     required: true
//     type: string
//
// produces:
//   - application/json
//
// responses:
//
//	'200':
//		description: Successful operation
//	'404':
//		description: Invalid recipe ID
func (h RecipesHandler) SearchRecipesHandler(c *gin.Context) {
	tag := c.Query("tag")
	cur, err := h.collection.Find(c.Request.Context(), bson.D{{Key: "tags", Value: tag}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cur.Close(c.Request.Context())
	recipes := make([]models.Recipe, 0)
	for cur.Next(c.Request.Context()) {
		var recipe models.Recipe
		err = cur.Decode(&recipe)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, recipes)
}
