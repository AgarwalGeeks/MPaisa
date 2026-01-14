package main

import (
	"database/sql"
	"fmt"
	"log"

	api "github.com/AgarwalGeeks/MPaisa/api"
	db "github.com/AgarwalGeeks/MPaisa/db/sqlc"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json") // Change to JSON
	viper.AddConfigPath(".")

	// Automatically override with environment variables if present
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, using environment variables if available: %s", err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("database migrated successfully")
}

func main() {
	fmt.Println("Hello, World!")

	initConfig()

	dbDriver := viper.GetString("db_driver")
	dbSource := viper.GetString("db_source")
	serverAddress := viper.GetString("server_address")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	migrationURL := viper.GetString("migration_url")
	runDBMigration(migrationURL, dbSource)

	store := db.NewStore(conn) // Replace with your actual store initialization
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	fmt.Println("Server is running on", serverAddress)
}
