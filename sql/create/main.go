package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// tag::schema[]
var schema = `CREATE TABLE sample (
	i INTEGER,
	s TEXT,
	t DATETIME DEFAULT CURRENT_TIMESTAMP
)`

// end::schema[]

// tag::opencreate[]
func main() {
	db, err := sql.Open("sqlite3", "file:testdb?mode=rwc")
	// end::opencreate[]
	if err != nil {
		panic(err)
	}
	// tag::close[]
	defer db.Close()
	// end::close[]
	// tag::exec[]
	_, err = db.Exec(schema)
	if err != nil {
		panic(err)
	}
	// end::exec[]
}
