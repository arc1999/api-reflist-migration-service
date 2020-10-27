package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type ReferenceListMongo struct {
	ID          int64     `json:"_id" bson:"_id"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	DateCreated time.Time `json:"dateCreated"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Status      int       `json:"status"`
	Slug primitive.Binary `json:"slug"`
}