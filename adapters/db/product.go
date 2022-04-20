package db

import (
	"database/sql"
	"go-hexagonal/application"

	_ "github.com/mattn/go-sqlite3"
)

// implements ProductPersistenceInterface
type ProductDB struct {
	db *sql.DB	
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{ db: db }
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

// Save(product ProductInterface) (ProductInterface, error)