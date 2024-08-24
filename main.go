package main

import (
	"database/sql"
	db2 "github.com/Jhon-Henkel/full_cycle_hexagonal_arc/adapters/db"
	"github.com/Jhon-Henkel/full_cycle_hexagonal_arc/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDB(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Test", 25)

	productService.Enable(product)
}
