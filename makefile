run:
	go run cmd/main.go
table:
	migrate create -dir internal/database/migration -ext sql db
table-up:
	migrate -path internal/database/migration -database "postgres://postgres:2005@localhost:5432/docker?sslmode=disable" up
table-down:
	migrate -path internal/database/migration -database "postgres://postgres:2005@localhost:5432/docker?sslmode=disable" down
force:
	migrate -path internal/database/migration -database "postgres://postgres:2005@localhost:5432/docker?sslmode=disable" force 20240802113108