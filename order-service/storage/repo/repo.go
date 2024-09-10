package repo

import orderspb "order-service/protos/order"

type OrderStoreI interface {
	CreateOrders(*orderspb.CreateOrdersRequest) (*orderspb.OrdersResponse, error)
	GetOrders(*orderspb.GetOrdersRequest) (*orderspb.OrdersResponse, error)
	UpdateOrders(*orderspb.UpdateOrdersRequest) (*orderspb.UpdateOrdersResponse,error)
	DeleteOrders(*orderspb.DeleteOrdersRequest) (*orderspb.DeleteOrdersResponse, error)
}
