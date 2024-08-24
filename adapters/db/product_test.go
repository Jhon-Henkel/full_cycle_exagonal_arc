package db_test

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")
	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (id string, name string, price float, status string);`
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products (id, name, price, status) values ("abc", "Product Test", 0, "disabled");`
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err.Error())
	}
	stmt.Exec()
}
