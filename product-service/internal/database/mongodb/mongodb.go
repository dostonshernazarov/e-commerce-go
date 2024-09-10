package mongodb

import (
	"context"
	"fmt"
	"product-service/internal/models"
	ppb "product-service/protos"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m MongoRepo) AddProductIntoMongo(ctx context.Context, req *ppb.ProductInfo) (*ppb.ProductID, error) {
	productCollection := m.Client.Database("products").Collection("products")
	req.CreatedAt =time.Now().Format(time.ANSIC)
	req.UpdatedAt =time.Now().Format(time.ANSIC)
	
	result, err := productCollection.InsertOne(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to insert products into mongoDB :%v", err)
	}
	id := result.InsertedID.(primitive.ObjectID).Hex()
	return &ppb.ProductID{Id: id}, nil
}

func (m MongoRepo) GetProductsFromMongo(req *ppb.Empty) ([]models.ProductWithID, error) {
	productCollection := m.Client.Database("products").Collection("products")

	cursor, err := productCollection.Find(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to find products from mongoDB :%v", err)
	}
	defer cursor.Close(context.Background())

	var products []models.ProductWithID
	cursor.All(context.Background(), products)
	return products, nil
}
