package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateHigherLevelCategoryTable() error {
	fmt.Println("Entered CreateHigherLevelCategoryTable")

	query := `create table if not exists higher_level_category (
		higher_level_category_id SERIAL PRIMARY KEY,
    	higher_level_category_name VARCHAR(100) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_by INT
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateHigherLevelCategoryTable")

	return err
}