package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Blog struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
}
