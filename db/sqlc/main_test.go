package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testStore   *Store
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://haagarwa:Harshit@12345@localhost:5435/MPaisa?sslmode=disable"
)

func TestMain(m *testing.M) {
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
