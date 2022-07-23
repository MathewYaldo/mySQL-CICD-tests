package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestMySQLDatabaseConnection(t *testing.T) {
	db, err := sql.Open("mysql", "user:password@/dbname")
// 	if err != nil {
// 		t.Fatal("A connection could not be established.")
// 	}
	if err = db.Ping(); err != nil {
	    t.Fatal("A connection could not be established.")
            d.Close()
        }
}
