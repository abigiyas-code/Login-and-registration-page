// In the main code we test the db file, which is the query does arrange the input properly
//	from the driver IP and the db source that we give from the query.sql.go file and
//    If it accepts save the given input as .sql, if not return error
package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root@:26257/users?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot import from the data base %v", err)
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}
