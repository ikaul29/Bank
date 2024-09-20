package main

import (
	"database/sql"
	"log"
	
    _ "github.com/lib/pq"
	"github.com/ikaul29/Bank/api"
	db "github.com/ikaul29/Bank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	//establish connection
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
    if err != nil {
		log.Fatal("cannot start server:", err)
	}
}