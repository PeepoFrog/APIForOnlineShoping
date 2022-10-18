package model

// id was primitive.ObjectID
type Commodity struct {
	ID       string  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Price    float64 `json:"price,omitempty"`
	ImageURL string  `json:"imageUrl,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
