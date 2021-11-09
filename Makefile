postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=20210918 -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root gobank

dropdb:
	docker exec -it postgres13 dropdb gobank

migrateup:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server