package types

import (
	"time"
)

type Category struct {
	ID                int       `json:"id"`
	Name              string    `json:"firstName"`
	Category          string    `json:"category"`
	Number			  int64     `json:"number"`
	NumberOfProducts   int       `json:"numberOfProducts"`  
	Quantity          int64     `json:"quantity"`
	CreatedAt         time.Time `json:"createdAt"`
}
