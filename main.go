package main

import (
	"database/sql"
	"log"

	"github.com/antunesluis/simplebank/api"
	db "github.com/antunesluis/simplebank/db/sqlc"
	"github.com/antunesluis/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource) // Use the package-level testDB variable
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
