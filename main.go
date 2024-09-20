package main

import (
	"database/sql"
	"log"

	"github.com/ikaul29/Bank/api"
	db "github.com/ikaul29/Bank/db/sqlc"
	"github.com/ikaul29/Bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	
	//establish connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
    if err != nil {
		log.Fatal("cannot start server:", err)
	}
}