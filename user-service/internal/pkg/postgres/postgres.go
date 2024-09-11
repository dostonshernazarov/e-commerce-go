package postgres

import (
	"database/sql"
	"fmt"
	"user-service/internal/pkg/load"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg load.Config) (*sql.DB, error) {
	target := fmt.Sprintf("host=%s port=%d dbname=%s password=%s sslmode=disable",
		cfg.Postgres.HOST, cfg.Postgres.PORT, cfg.Postgres.DBNAME, cfg.Postgres.PASSWORD)
	db, err := sql.Open("postgres", target)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
