package types

import (
	"time"
)

type Item struct {
	ID                int       `json:"id"`
	Name              string    `json:"firstName"`
	Price			  float64 	`json:"price"`
	Store_ID		  int		`json:"store_id"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int     `json:"stock_quantity"`
	Created_At        time.Time `json:"created_at"`
	Created_By		  int		`json:"created_by"`
}

type Create_Item struct {
	Name              string    `json:"firstName"`
	Price			  float64 	`json:"price"`
	Store_ID		  int		`json:"store_id"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int    `json:"stock_quantity"`
}

type Update_Item struct {
	ID				  int       `json:"id"`
	Name              string    `json:"firstName"`
	Price			  float64 	`json:"price"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int     `json:"stock_quantity"`
}

type Delete_Item struct {
	ID		int `json:"id"`
}

func New_Item(name string, price float64, category_id int, store_id int, stock_quantity int)(*Item, error) {
	return &Item{
	Name:           name,
	Price:          price,
	Category_ID:    category_id,
	Stock_Quantity: stock_quantity,
	Store_ID: store_id,
}, nil
}

func New_Update_Item(id int, name string, price float64, category_id int, stock_quantity int)(*Update_Item, error) {
	return &Update_Item{
	ID:             id,
	Name:           name,
	Price:          price,
	Category_ID:    category_id,
	Stock_Quantity: stock_quantity,
}, nil
}
