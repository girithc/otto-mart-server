package storage
import (
	"pronto-go/types"
	_ "github.com/lib/pq"
)

func (s *PostgresStore) CreateProductTable() error {
	//fmt.Println("Entered CreateProductTable -- product.go")

	query := `create table if not exists product (
		id serial primary key,
		name varchar(100),
		category varchar(100),
		number serial,
		quantity int,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)

	//fmt.Println("Exiting CreateProductTable -- product.go")

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