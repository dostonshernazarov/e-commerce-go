package storage

import (
	"database/sql"
	"order-service/storage/postgres"
	"order-service/storage/repo"
)

type IStorage interface {
	Order() repo.OrderStoreI
}

type storagePg struct {
	db       *sql.DB
	OrderRepo repo.OrderStoreI
}

func NewStoragePg(db *sql.DB) *storagePg {
	return &storagePg{
		db:       db,
		OrderRepo: postgres.NewOrderRepo(db),
	}
}

func (s storagePg) Order() repo.OrderStoreI {
	return s.OrderRepo
}