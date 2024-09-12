run:
	@go run ./cmd/main.go
migrate:
	migrate create -dir ./storage/db -ext sql db

migrate-up:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/comment?sslmode=disable" up

migrate-down:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/comment?sslmode=disable" down

migrate-force:
	migrate -path ./storage/db -database "postgres://postgres:dilshod@localhost:5432/comment?sslmode=disable" force 20240909105200
