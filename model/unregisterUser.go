package model

type UnregUser struct {
	ID     string `json:"_id,omitempty" bson:"_id,omitempty"`
	Basket Basket `json:"basket,omitempty"`
	Name   string `json:"name,omitempty"`
}
