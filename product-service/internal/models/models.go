package models

type ProductWithID struct {
	ID           string            `json:"id" bson:"id"`
	Name         string            `json:"name" bson:"name"`
	CategoryName string            `json:"category_name" bson:"category_name"`
	Quantity     int32             `json:"quantity" bson:"quantity"`
	Price        float32           `json:"price" bson:"price"`
	Created_at string          `json:"created_at" bson:"created_at"`
	Updated_at string          `json:"updated_at" bson:"updated_at"`
	Details      map[string]string `json:"details" bson:"details"`
}
