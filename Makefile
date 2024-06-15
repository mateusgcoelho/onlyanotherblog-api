run:
	go run ./cmd/api/main.go

migrate-up:
	migrate -path ./database/migration/ -database postgres://postgres:Docker@localhost:5432/onlyanotherblog?sslmode=disable up

migrate-down:
	migrate -path ./database/migration/ -database postgres://postgres:Docker@localhost:5432/onlyanotherblog?sslmode=disable down

generate-schemas:
	sqlc generate
