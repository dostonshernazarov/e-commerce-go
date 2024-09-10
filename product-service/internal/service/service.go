package service

import (
	"context"
	"fmt"
	"product-service/internal/database/mongodb"
	ppb "product-service/protos"
)

type ProductService struct {
	ppb.UnimplementedProductServiceServer
	Mongo *mongodb.MongoRepo
}

func NewProductService(m *mongodb.MongoRepo)*ProductService{
	return &ProductService{Mongo: m}
}

func (p *ProductService) CreateProduct(ctx context.Context, req *ppb.ProductInfo)(*ppb.ProductID, error){
	fmt.Println("start...")
	productID, err :=p.Mongo.AddProductIntoMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return productID, nil
}

func (p *ProductService) GetProducts(req *ppb.Empty, stream ppb.ProductService_GetProductsServer)( error){
	listProducts, err :=p.Mongo.GetProductsFromMongo(req)
	if err != nil {
		return err
	}
	for _, v :=range listProducts {
		if err :=stream.Send(&ppb.ProductWithID{
			Id: v.ID,
			Name: v.Name,
			CategoryName: v.CategoryName,
			Price: v.Price,
			Quantity: v.Quantity,
			CreatedAt: v.Created_at,
			UpdatedAt: v.Updated_at,
			Details: v.Details,	
		}); err != nil {
			return fmt.Errorf("unable to send product information")
		}
	}
	return nil
}

func (p *ProductService) GetProductByID(ctx context.Context, req *ppb.ProductID)(*ppb.ProductWithID, error){
	return nil, nil
}

func (p *ProductService) UpdateProductByID(ctx context.Context, req *ppb.ProductInfo)(*ppb.ProductResponse, error){
	return nil, nil
}

func (p *ProductService) DeleteProductByID(ctx context.Context, req *ppb.ProductID)(*ppb.ProductResponse, error){
	return nil, nil

}

