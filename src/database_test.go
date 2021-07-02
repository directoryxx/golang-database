package golangdatabase

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func TestEmpty(t *testing.T) {
	db,err := sql.Open("mysql","golang:golang@(database:3306)/golang")
	if err != nil {
		panic(err)
	}


	defer db.Close()
}
