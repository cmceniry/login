package main

import (
	"database/sql"
	"fmt"
	"time"

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
	rows, err := db.Query(
		"SELECT i, s, t FROM sample WHERE DATE(t) <= DATE(?)",
		time.Now(),
	)
	// end::query[]
	if err != nil {
		panic(err)
	}
	// tag::stmtclose[]
	defer rows.Close()
	// end::stmtclose[]
	// tag::next[]
	for rows.Next() {
		// end::next[]
		// tag::scan[]
		var i int64
		var s string
		var t time.Time
		err := rows.Scan(&i, &s, &t)
		// end::scan[]
		if err != nil {
			panic(err)
		}
		// tag::printout[]
		fmt.Printf("%s: %d %s\n", t, i, s)
		// end::printout[]
	}
}
