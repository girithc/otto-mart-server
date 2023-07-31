package storage

import (
	"fmt"
    "database/sql"
	_ "github.com/lib/pq"
)


type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	fmt.Println("Entered NewPostgresStore() -- db.go")

	connStr := "user=girithc dbname=prontodb password=password sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Exiting (Err) NewPostgresStore() -- db.go")
	
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Exiting (db.Ping()) NewPostgresStore() -- db.go")
	
		return nil, err
	}

	fmt.Println("Exiting NewPostgresStore() -- db.go")
	

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	fmt.Println("Entered Init() -- db.go")

	errProduct := s.CreateProductTable()
	if errProduct != nil{
		return errProduct
	} else {
		fmt.Println("	CreateProductTable success")
	}

	errCategory := s.CreateCategoryTable()
	if errCategory != nil{
		return errCategory
	} else {
		fmt.Println("	CreateCategoryTable success")
	}

	fmt.Println("Exiting Init() -- db.go")
	return nil
}

