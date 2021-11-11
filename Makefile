postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=20210918 -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root gobank

dropdb:
	docker exec -it postgres13 dropdb gobank

migration:
	migrate create -ext sql -dir db/migration -seq add_users

migrateup:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose up

migrateup1:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose down

migratedown1:
	migrate -path ./db/migration -database "postgresql://root:20210918@localhost:5432/gobank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/jdiego0102/gobank/db/sqlc Store

.PHONY: postgres createdb dropdb migration migrateup migratedown migrateup1 migratedown2 sqlc test server mock