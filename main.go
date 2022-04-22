package main

import (
	"database/sql"
	dbAdapter "go-hexagonal/adapters/db"
	"go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")

	productDBAdapter := dbAdapter.NewProductDB(db)
	productService := application.NewProductService(productDBAdapter)

	product, _ := productService.Create("Product 1", 10.0)

	productService.Enable(product)
}