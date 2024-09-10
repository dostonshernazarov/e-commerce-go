package postgres

import (
	orderspb "order-service/protos/order"
	"time"
)

func (o *OrderRepo) CreateOrders(req *orderspb.CreateOrdersRequest) (*orderspb.OrdersResponse, error) {
	// id, err := user.UserClinet(int(req.UserId))
	// if err != nil {
	// 	if id == 0 {
	// 		return nil, fmt.Errorf("not found user")
	// 	}
	// 	return nil, err
	// }

	// for _, p := range req.Products {
	// 	id, err := product.ProductClinet(p)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if id == 0 {
	// 		return nil, fmt.Errorf("not found product")
	// 	}
	// }

	res := &orderspb.Orders{}
	var created_at, updated_at time.Time
	err := o.db.QueryRow("insert into orders(user_id, location,status) values ($1,$2,$3) returning id, created_at, updated_at", req.UserId, req.Location, req.Status).Scan(
		&res.Id,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	res.UserId = req.UserId
	res.Products = req.Products
	res.CreatedAt = created_at.Format(time.RFC3339)
	res.UpdatedAt = updated_at.Format(time.RFC3339)

	for _, i := range req.Products {
		_, err := o.db.Exec("insert into order_products(order_id,product_id) values ($1,$2)", res.Id, i)
		if err != nil {
			return nil, err
		}
	}

	return &orderspb.OrdersResponse{Orders: res}, nil
}

func (o *OrderRepo) GetOrders(req *orderspb.GetOrdersRequest) (*orderspb.OrdersResponse, error) {
	rows, err := o.db.Query("select id, user_id, location,status, created_at, updated_at from orders where user_id = $1", req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := &orderspb.Orders{}
	for rows.Next() {
		var created_at, updated_at time.Time
		err := rows.Scan(
			&res.Id,
			&res.UserId,
			&res.Location,
			&res.Status,
			&created_at,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		res.CreatedAt = created_at.Format(time.RFC3339)
		res.UpdatedAt = updated_at.Format(time.RFC3339)
	}
	rows2, err := o.db.Query("select product_id from order_products where order_id = $1", res.Id)
	if err != nil {
		return nil, err
	}
	defer rows2.Close()
	res.Products = make([]string, 0)
	for rows2.Next() {
		var productId string
		err := rows2.Scan(&productId)
		if err != nil {
			return nil, err
		}
		res.Products = append(res.Products, productId)
	}
	return &orderspb.OrdersResponse{Orders: res}, nil
}

func (o *OrderRepo) UpdateOrders(req *orderspb.UpdateOrdersRequest) (*orderspb.UpdateOrdersResponse, error) {
	currentTime := time.Now()

	_, err := o.db.Exec(
		"UPDATE orders SET location = $1, status = $2, updated_at = $3 WHERE id = $4 AND user_id = $5",
		req.Location, req.Status, currentTime, req.Id, req.UserId,
	)
	if err != nil {
		return nil, err
	}

	_, err = o.db.Exec("DELETE FROM order_products WHERE order_id = $1", req.Id)
	if err != nil {
		return nil, err
	}

	for _, productId := range req.Products {
		_, err := o.db.Exec("INSERT INTO order_products (order_id, product_id) VALUES ($1, $2)", req.Id, productId)
		if err != nil {
			return nil, err
		}
	}

	return &orderspb.UpdateOrdersResponse{
		Message: "Order updated successfully",
	}, nil
}

func (o *OrderRepo) DeleteOrders(req *orderspb.DeleteOrdersRequest) (*orderspb.DeleteOrdersResponse, error) {
	_, err := o.db.Exec("delete from orders where id = $1", req.Id)
	if err != nil {
		return nil, err
	}
	return &orderspb.DeleteOrdersResponse{
		Message: "Order deleted successfully",
	}, nil
}
