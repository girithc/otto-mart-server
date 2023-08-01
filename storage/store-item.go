package storage

import (
	"fmt"
	"pronto-go/types"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateItemTable() error {
	fmt.Println("Entered CreateItemTable")

	query := `create table if not exists item (
		item_id SERIAL PRIMARY KEY,
		item_name VARCHAR(100) NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		store_id INT REFERENCES Store(store_id) ON DELETE CASCADE,
		category_id INT REFERENCES Category(category_id) ON DELETE CASCADE,
		stock_quantity INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_by INT
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateItemTable")

	return err
}

func (s *PostgresStore) CreateProduct(p *types.Product) error {
	query := `insert into product 
	(name, category, number, quantity, created_at)
	values ($1, $2, $3, $4, $5)`
	_, err := s.db.Query(
		query,
		p.Name,
		p.Category,
		p.Number,
		p.Quantity,
		p.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}