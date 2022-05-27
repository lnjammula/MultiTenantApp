postgres:
	docker run --name postgres14 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.3-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root MultiTenantApp 

dropdb:
	docker exec -it postgres14 dropdb MultiTenantApp

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5433/MultiTenantApp?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5433/MultiTenantApp?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server