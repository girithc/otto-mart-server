package storage

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateCategoryHigherLevelMappingTable() error {
	fmt.Println("Entered CreateCategoryHigherLevelMappingTable")

	query := `create table if not exists category_higher_level_mapping (
		category_higher_level_mapping_id SERIAL PRIMARY KEY,
		higher_level_category_id INT REFERENCES Higher_Level_Category(higher_level_category_id) ON DELETE CASCADE,
		category_id INT REFERENCES Category(category_id) ON DELETE CASCADE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		created_by INT
	)`

	_, err := s.db.Exec(query)

	fmt.Println("Exiting CreateCategoryHigherLevelMappingTable")

	return err
}