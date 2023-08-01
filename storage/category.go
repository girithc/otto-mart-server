package storage

import (
	"pronto-go/types"
	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateCategoryTable() error {
	//fmt.Println("Entered CreateCategoryTable -- category.go")

	query := `create table if not exists category (
		id serial primary key,
		name varchar(100),
		parent_category bool,
		number serial unique,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	//fmt.Println("Exiting CreateCategoryTable -- category.go")

	return err
}

func (s *PostgresStore) CreateCategory(c *types.Category) error {
	query := `insert into category 
	(name, parent_category, number, created_at)
	values ($1, $2, $3, $4)`
	_, err := s.db.Query(
		query,
		c.Name,
		c.ParentCategory,
		c.Number,
		c.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}