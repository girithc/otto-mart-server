package storage

import (
	"database/sql"
	"fmt"
	"pronto-go/types"

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

func (s *PostgresStore) Create_Store(st *types.Store) (*types.Store, error) {
	query := `insert into store
	(store_name, address, created_by) 
	values ($1, $2, $3) returning store_id, store_name, address, created_at, created_by
	`
	rows, err := s.db.Query(
		query,
		st.Store_Name,
		st.Address,
		st.Created_By)

	fmt.Println("CheckPoint 1")

	if err != nil {
		return nil, err
	}

	fmt.Println("CheckPoint 2")

	stores := []*types.Store{}
	
	for rows.Next() {
		store, err := scan_Into_Store(rows)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	fmt.Println("CheckPoint 3")

	return stores[0], nil
}

func (s *PostgresStore) Get_Stores() ([]*types.Store, error) {
	rows, err := s.db.Query("select * from store")

	if err != nil {
		return nil, err
	}

	stores := []*types.Store{}
	for rows.Next() {
		store, err := scan_Into_Store(rows)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}

func (s *PostgresStore) Get_Store_By_ID(store_id int) (*types.Store, error) {
	row, err := s.db.Query("select * from store where store_id = $1", store_id)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		return scan_Into_Store(row)
	}

	return nil, fmt.Errorf("store with id = [%d] not found", store_id)
}

func (s *PostgresStore) Update_Store(st *types.Update_Store) (*types.Update_Store,error) {
	query := `update store
	set store_name = $1, address = $2
	where store_id = $3 
	returning store_name, address, store_id`
	
	rows, err := s.db.Query(
		query, 
		st.Store_Name,
		st.Address,
		st.Store_ID,
	)

	if err != nil {
		return nil, err
	}

	stores := []*types.Update_Store{}
	
	for rows.Next() {
		store, err := scan_Into_Update_Store(rows)
		if err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}
	

	return stores[0], nil
}

func (s *PostgresStore) Delete_Store(store_id int) error {
	_, err := s.db.Query("delete from store where store_id = $1", store_id)
	return err
}

func scan_Into_Store (rows *sql.Rows) (*types.Store, error) {
	store := new(types.Store)
	err := rows.Scan(
		&store.Store_ID,
		&store.Store_Name,
		&store.Address,
		&store.Created_At,
		&store.Created_By,
	)

	return store, err
}

func scan_Into_Update_Store(rows *sql.Rows) (*types.Update_Store, error) {
	store := new(types.Update_Store)
	error := rows.Scan(
		&store.Store_Name,
		&store.Address,
		&store.Store_ID,
	)

	return store, error
} 