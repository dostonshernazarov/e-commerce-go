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
	product, err :=p.Mongo.GetProductByIDFromMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductService) UpdateProductByID(ctx context.Context, req *ppb.ProductInfo)(*ppb.ProductResponse, error){
	response, err :=p.Mongo.UpdateProductByIDFromMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p *ProductService) DeleteProductByID(ctx context.Context, req *ppb.ProductID)(*ppb.ProductResponse, error){
	response, err :=p.Mongo.DeleteProductByIDFromMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func (p ProductService) UpdateProductQuantity(ctx context.Context, req *ppb.QuantityRequest)(*ppb.ProductResponse,error){
	response, err :=p.Mongo.UpdateProductQuantityFromMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (p ProductService) GetProductByCategory( ctx context.Context, req *ppb.CategoryName)(*ppb.ListProducts, error){
	listProducts, err :=p.Mongo.GetProductByCategoryFromMongo(ctx, req)
	if err != nil {
		return nil, err
	}
	return listProducts, nil
}