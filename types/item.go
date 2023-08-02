package types

import (
	"time"
)

type Product struct {
	ID                int       `json:"id"`
	Name              string    `json:"firstName"`
	Category          string    `json:"category"`
	Number            int64     `json:"number"`
	Quantity          int64     `json:"quantity"`
	CreatedAt         time.Time `json:"createdAt"`
}
