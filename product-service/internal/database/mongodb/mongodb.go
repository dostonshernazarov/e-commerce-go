package mongodb

import (
	"context"
	"fmt"
	"product-service/internal/models"
	ppb "product-service/protos"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (m MongoRepo) AddProductIntoMongo(ctx context.Context, req *ppb.ProductInfo) (*ppb.ProductID, error) {
	productCollection := m.Client.Database("products").Collection("products")
	req.CreatedAt = time.Now().Format(time.ANSIC)
	req.UpdatedAt = time.Now().Format(time.ANSIC)

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

func (m MongoRepo) GetProductByIDFromMongo(ctx context.Context, req *ppb.ProductID) (*ppb.ProductWithID, error) {
	productCollection := m.Client.Database("products").Collection("products")

	idObj, _ := primitive.ObjectIDFromHex(req.Id)

	result := productCollection.FindOne(ctx, bson.M{"_id": idObj})
	var product models.ProductWithID
	if err := result.Decode(&product); err != nil {
		return nil, fmt.Errorf("failed to decode product in getByID : %v", err)
	}

	return &ppb.ProductWithID{
		Id:           req.Id,
		Name:         product.Name,
		CategoryName: product.CategoryName,
		Price:        product.Price,
		Quantity:     product.Quantity,
		CreatedAt:    product.Created_at,
		UpdatedAt:    product.Updated_at,
		Details:      product.Details,
		Images:       product.Images,
	}, nil
}

func (m MongoRepo) UpdateProductByIDFromMongo(ctx context.Context, req *ppb.ProductInfo) (*ppb.ProductResponse, error) {
	productCollection := m.Client.Database("products").Collection("products")

	idObj, _ := primitive.ObjectIDFromHex(ctx.Value("id").(string))
	result, err := productCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{
		"name":      req.Name,
		"quantity":  req.Quantity,
		"price":     req.Price,
		"updatedat": time.Now().Format(time.ANSIC),
		"details":   req.Details,
		"images":    req.Images,
	}})
	if err != nil {
		return nil, fmt.Errorf("failed to update product info : %v", err)
	}
	if result.ModifiedCount == 0 {
		return nil, fmt.Errorf("no product found with ID: %v", ctx.Value("id").(string))
	}
	return &ppb.ProductResponse{Message: "Product updated succesfully"}, nil
}

func (m MongoRepo) DeleteProductByIDFromMongo(ctx context.Context, req *ppb.ProductID) (*ppb.ProductResponse, error) {
	productCollection := m.Client.Database("products").Collection("products")

	idObj, _ := primitive.ObjectIDFromHex(req.Id)
	result, err := productCollection.DeleteOne(ctx, bson.M{"_id": idObj})
	if err != nil {
		return nil, fmt.Errorf("failed to delete product info : %v", err)
	}
	if result.DeletedCount == 0 {
		return nil, fmt.Errorf("no product found with ID: %v", req.Id)
	}
	return &ppb.ProductResponse{Message: "Product  Deleted succesfully"}, nil
}
func (m MongoRepo) UpdateProductQuantityFromMongo(ctx context.Context, req *ppb.QuantityRequest) (*ppb.ProductResponse, error) {
	productCollection := m.Client.Database("products").Collection("products")

	idObj, _ := primitive.ObjectIDFromHex(req.Id)

	projection := bson.D{{"quantity", 1}, {"_id", 0}}

	result := productCollection.FindOne(ctx, bson.M{"_id": idObj}, options.FindOne().SetProjection(projection))
	var quantity models.ProductQuantity
	if err := result.Decode(&quantity); err != nil {
		return nil, fmt.Errorf("failed to retrive quantity from mongoDB: %v", err)
	}

	_, err := productCollection.UpdateOne(ctx, bson.M{"_id": idObj}, bson.M{"$set": bson.M{
		"quantity": quantity.Quantity - req.Quantity,
	}})
	if err != nil {
		return nil, fmt.Errorf("failed to update quantity from mongoDB: %v", err)
	}
	return &ppb.ProductResponse{Message: "Product Quantity updated succesfully"}, nil
}

func (m MongoRepo) GetProductByCategoryFromMongo(ctx context.Context, req *ppb.CategoryName) (*ppb.ListProducts, error) {
	productCollection := m.Client.Database("products").Collection("products")

	cursor, err := productCollection.Find(ctx, bson.M{"categoryname": req.Name})
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category : %v", err)
	}
	defer cursor.Close(ctx)

	var list []models.ProductWithMongoID
	if err := cursor.All(ctx, &list); err != nil {
		return nil, fmt.Errorf("failed to cursor all  products by category : %v", err)
	}

	var listProducts ppb.ListProducts
	for i, v := range list {
		listProducts.Listproducts = append(listProducts.Listproducts, &ppb.ProductWithID{
			Id:           list[i].ID.Hex(),
			Name:         v.Name,
			CategoryName: v.CategoryName,
			Price:        v.Price,
			Quantity:     v.Quantity,
			CreatedAt:    v.Created_at,
			UpdatedAt:    v.Updated_at,
			Details:      v.Details,
			Images:       v.Images,
		})
	}
	return &listProducts, nil
}
