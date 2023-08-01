package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateStoreTable() error {
	fmt.Println("Entered CreateStoreTable")

	query := `create table if not exists store (
		store_id SERIAL PRIMARY KEY,
		store_name VARCHAR(100) NOT NULL,
		address VARCHAR(200) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_by INT	
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateStoreTable")

	return err
}