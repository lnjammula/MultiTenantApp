package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"multitenant.com/app/api"
	db "multitenant.com/app/db/sqlc"
	"multitenant.com/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannnot connect to DB", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
