package service

import (
	"context"
	"database/sql"
	orderpb "order-service/protos/order"
	"order-service/storage"
)

type OrderService struct {
	storage storage.IStorage
	orderpb.UnimplementedOrdersServiceServer
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		storage: storage.NewStoragePg(db),
	}
}

func (s *OrderService) CreateOrders(ctx context.Context, req *orderpb.CreateOrdersRequest) (*orderpb.OrdersResponse, error) {
	res, err := s.storage.Order().CreateOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) GetOrders(ctx context.Context, req *orderpb.GetOrdersRequest) (*orderpb.OrdersResponse, error) {
	res, err := s.storage.Order().GetOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) UpdateOrders(ctx context.Context, req *orderpb.UpdateOrdersRequest) (*orderpb.UpdateOrdersResponse, error) {
	res, err := s.storage.Order().UpdateOrders(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *OrderService) DeleteOrders(ctx context.Context, req *orderpb.DeleteOrdersRequest) (*orderpb.DeleteOrdersResponse, error) {
	res, err := s.storage.Order().DeleteOrders(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}