package db_test

import (
	"database/sql"
	"go-hexagonal/adapters/db"
	"go-hexagonal/application"
	"testing"

	"github.com/stretchr/testify/require"
)

var DB *sql.DB

func setup() {
	DB, _ = sql.Open("sqlite3", ":memory:")

	createTable(DB)
	createProduct(DB)
}

func createTable(db *sql.DB) {
	createTableCmd := `CREATE TABLE products (id STRING, name STRING, price FLOAT, status STRING);`

	stmt, err := db.Prepare(createTableCmd)
	
	if err != nil {
		panic(err)
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insertCmd := `INSERT INTO products VALUES (?, ?, ?, ?)`

	stmt, err := db.Prepare(insertCmd)

	if err != nil {
		panic(err)
	}

	stmt.Exec("1", "Product 1", 10.0, "disabled")
}

func TestProductDB_Get(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product, err := productDB.Get("1")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDB_Save(t *testing.T) {
	setup()
	defer DB.Close()

	productDB := db.NewProductDB(DB)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 20.0

	productResult, err := productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, "Product 2", productResult.GetName())
	require.Equal(t, 20.0, productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDB.Save(product)

	require.Nil(t, err)
	require.Equal(t, "Product 2", productResult.GetName())
	require.Equal(t, 20.0, productResult.GetPrice())
	require.Equal(t, "enabled", productResult.GetStatus())
}