package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// tag::open[]
func main() {
	db, err := sql.Open("sqlite3", "file:testdb?mode=rw")
	// end::open[]
	if err != nil {
		panic(err)
	}
	// tag::close[]
	defer db.Close()
	// end::close[]
	// tag::query[]
	res, err := db.Exec(
		"INSERT INTO sample (i, s) VALUES (?, ?)",
		2,
		"2",
	)
	// end::query[]
	if err != nil {
		panic(err)
	}
	// tag::rows[]
	affected, err := res.RowsAffected()
	// end::rows[]
	if err != nil {
		panic(err)
	}
	// tag::print[]
	fmt.Printf("%d row(s) inserted\n", affected)
	// end::print[]
}
