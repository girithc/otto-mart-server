package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateCartItemTable() error {
	fmt.Println("Entered CreateCartItemTable")

	query := `create table if not exists cart_item (
		cart_item_id SERIAL PRIMARY KEY,
		cart_id INT REFERENCES Shopping_Cart(cart_id) ON DELETE CASCADE,
		item_id INT REFERENCES Item(item_id) ON DELETE CASCADE,
		quantity INT NOT NULL 
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateCartItemTable")

	return err
}