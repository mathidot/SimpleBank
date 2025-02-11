package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSourceName = "postgresql://root:secret@172.23.206.114:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var connPool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	connPool, err = pgxpool.New(context.Background(), dbSourceName)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(connPool)
	fmt.Println("Connect to postgres successfully")
	os.Exit(m.Run())
}
