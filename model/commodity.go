package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Commodity struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"movie,omitempty"`
	Price    float64            `json:"price,omitempty"`
	ImageURL string             `json:"imageUrl,omitempty"`
	Quantity int                `json:"quantity,omitempty"`
}
