package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	testQueries *Queries
	testStore   *Store
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")  // Change to JSON
	viper.AddConfigPath("../..") // Adjust path to locate config.json

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func TestMain(m *testing.M) {
	initConfig()

	dbDriver := viper.GetString("db_driver")
	dbSource := viper.GetString("db_source")

	// code to run before tests
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)
	testStore = NewStore(conn)

	// run the tests
	os.Exit(m.Run())

	// code to run after tests
}
