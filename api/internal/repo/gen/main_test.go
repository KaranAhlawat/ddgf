package repo

import (
	"log"
	"os"
	"testing"

	repo "github.com/KaranAhlawat/ddgf/internal/repo/postgresql"
	_ "github.com/lib/pq"
)

var test_queries *Queries

func TestMain(m *testing.M) {
	conn, err := repo.InitPostgresConn("test")
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	test_queries = New(conn)

	os.Exit(m.Run())
}
