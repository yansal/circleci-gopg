package main

import (
	"database/sql"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

func TestMain(t *testing.T) {
	t.Log(os.Getenv("DATABASE_URL"))
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	row := db.QueryRow("select now();")
	var now time.Time
	if err := row.Scan(&now); err != nil {
		t.Fatal(err)
	}
	t.Log(now)
}
