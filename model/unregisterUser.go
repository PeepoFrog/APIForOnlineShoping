package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UnregUser struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Basket Basket             `json:"basket,omitempty"`
}
