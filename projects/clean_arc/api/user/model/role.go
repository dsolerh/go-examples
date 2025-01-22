package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleCode string

const (
	RoleCodeLearner RoleCode = "LEARNER"
	RoleCodeAdmin   RoleCode = "ADMIN"
	RoleCodeAuthor  RoleCode = "AUTHOR"
	RoleCodeEditor  RoleCode = "EDITOR"
)

const RolesCollectionName = "roles"

type Role struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Code      RoleCode           `bson:"code" validate:"required,rolecode"`
	Status    bool               `bson:"status" validate:"required"`
	CreatedAt time.Time          `bson:"createdAt" validate:"required"`
	UpdatedAt time.Time          `bson:"updatedAt" validate:"required"`
}
