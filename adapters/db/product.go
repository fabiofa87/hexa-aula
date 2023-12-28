package db

import (
	"database/sql"

	"github.com/go-hexa/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func (p *ProductDB) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := p.db.Prepare("SELECT id, name, status, price FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
