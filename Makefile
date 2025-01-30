run:
	go run cmd/api/main.go

test:
	go test -v ./...

migrate:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/rental?sslmode=disable" up

docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-down:
	docker-compose down