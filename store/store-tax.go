package store

import (
	"database/sql"
	"fmt"
)

func (s *PostgresStore) CreateTaxTable(tx *sql.Tx) error {
	// SQL query to create the 'tax' table with INTEGER fields and defaults
	createTaxTableQuery := `
    CREATE TABLE IF NOT EXISTS tax (
        id SERIAL PRIMARY KEY,
        sgst INTEGER NOT NULL DEFAULT 0 CHECK (sgst >= 0 AND sgst <= 10000),
        cgst INTEGER NOT NULL DEFAULT 0 CHECK (cgst >= 0 AND cgst <= 10000),
        gst INTEGER NOT NULL DEFAULT 0 CHECK (gst >= 0 AND gst <= 10000),
        cess INTEGER NOT NULL DEFAULT 0 CHECK (cess >= 0 AND cess <= 10000),
        total_tax INTEGER GENERATED ALWAYS AS (gst + cess) STORED
    );`

	_, err := tx.Exec(createTaxTableQuery)
	if err != nil {
		return fmt.Errorf("error creating tax table: %w", err)
	}

	return nil
}

func (s *PostgresStore) GetTaxDetails() ([]TaxDetail, error) {
	var taxDetails []TaxDetail

	query := `SELECT id, gst, cess FROM tax`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing tax details query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var taxDetail TaxDetail
		if err := rows.Scan(&taxDetail.ID, &taxDetail.GST, &taxDetail.CESS); err != nil {
			return nil, fmt.Errorf("error scanning tax row: %w", err)
		}
		taxDetails = append(taxDetails, taxDetail)
	}

	// Check for any error that might have occurred during iteration
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating tax rows: %w", err)
	}

	return taxDetails, nil
}

type TaxDetail struct {
	ID   int
	GST  int
	CESS int
}
