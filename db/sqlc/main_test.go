package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/jdiego0102/gobank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("No hay conexión a la base de datos: ", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
