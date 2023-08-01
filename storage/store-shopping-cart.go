package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateShoppingCartTable() error {
	fmt.Println("Entered CreateShoppingCartTable")

	query := `create table if not exists shopping_cart (
		cart_id SERIAL PRIMARY KEY,
    	customer_id INT REFERENCES Customer(customer_id) ON DELETE CASCADE
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateShoppingCartTable")

	return err
}