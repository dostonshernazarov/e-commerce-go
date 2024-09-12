package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductWithID struct {
	ID           string            `json:"id" bson:"id"`
	Name         string            `json:"name" bson:"name"`
	CategoryName string            `json:"category_name" bson:"category_name"`
	Quantity     int32             `json:"quantity" bson:"quantity"`
	Price        float32           `json:"price" bson:"price"`
	Created_at   string            `json:"created_at" bson:"created_at"`
	Updated_at   string            `json:"updated_at" bson:"updated_at"`
	Details      map[string]string `json:"details" bson:"details"`
	Images       []string          `json:"images" bson:"images"`
}

type ProductWithMongoID struct {
	ID           primitive.ObjectID  `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	CategoryName string             `json:"category_name" bson:"category_name"`
	Quantity     int32              `json:"quantity" bson:"quantity"`
	Price        float32            `json:"price" bson:"price"`
	Created_at   string             `json:"created_at" bson:"created_at"`
	Updated_at   string             `json:"updated_at" bson:"updated_at"`
	Details      map[string]string  `json:"details" bson:"details"`
	Images       []string           `json:"images" bson:"images"`
}

type ProductInfo struct {
	Name         string            `json:"name" bson:"name"`
	CategoryName string            `json:"category_name" bson:"category_name"`
	Quantity     int32             `json:"quantity" bson:"quantity"`
	Price        float32           `json:"price" bson:"price"`
	Created_at   string            `json:"created_at" bson:"created_at"`
	Updated_at   string            `json:"updated_at" bson:"updated_at"`
	Details      map[string]string `json:"details" bson:"details"`
	Images       []string          `json:"images" bson:"images"`
}
type ProductQuantity struct {
	Quantity int32 `json:"quantity" bson:"quantity"`
}

type CategoryName struct {
	Name string `json:"name" bson:"name"`
}
