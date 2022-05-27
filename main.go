package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"multitenant.com/app/api"
	db "multitenant.com/app/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://root:secret@localhost:5433/MultiTenantApp?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannnot connect to DB", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
