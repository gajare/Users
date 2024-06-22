package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID `json:"id,omitempty" bsdon:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Price       float64            `json:"price" bson:"price"`
}
