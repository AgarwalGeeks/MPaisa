package main

import (
	"database/sql"
	"fmt"
	"log"

	api "github.com/AgarwalGeeks/MPaisa/api"
	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://haagarwa:Harshit@12345@localhost:5435/MPaisa?sslmode=disable"
	serverAddress = "0.0.0.0:8090"
)

func main() {
	fmt.Println("Hello, World!")

	// code to run before tests
	conn, err := sql.Open(dbDriver, dbSource)

	store := db.NewStore(conn) // Replace with your actual store initialization
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		panic("cannot start server: " + err.Error())
	}

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	fmt.Println("Server is running on", serverAddress)
}
