package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://ulule:ulule@localhost:35432/ulule?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("select now();")
	var now time.Time
	if err := row.Scan(&now); err != nil {
		log.Fatal(err)
	}
	fmt.Println(now)
}
