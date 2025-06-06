package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Recipe struct {
	ID           bson.ObjectID `json:"id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	Tags         []string      `json:"tags" bson:"tags"`
	Ingredients  []string      `json:"ingredients" bson:"ingredients"`
	Instructions []string      `json:"instructions" bson:"instructions"`
	PublishedAt  time.Time     `json:"published_at" bson:"published_at"`
}
