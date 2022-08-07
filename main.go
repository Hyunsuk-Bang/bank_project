package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // for postgres
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:admin@localhost:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
