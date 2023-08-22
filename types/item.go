package types

import (
	"time"
)

//Basic
type Item struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Price			  float64 	`json:"price"`
	Store_ID		  int		`json:"store_id"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int     `json:"stock_quantity"`
	Created_At        time.Time `json:"created_at"`
	Created_By		  int		`json:"created_by"`
}

type Create_Item struct {
	Name              string    `json:"name"`
	Price			  float64 	`json:"price"`
	Store_ID		  int		`json:"store_id"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int    `json:"stock_quantity"`
}

type Update_Item struct {
	ID				  int       `json:"id"`
	Name              string    `json:"name"`
	Price			  float64 	`json:"price"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int     `json:"stock_quantity"`
}

type Delete_Item struct {
	ID		int `json:"id"`
}

//Custom
type Get_Items_By_CategoryID_And_StoreID struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	Price			  float64 	`json:"price"`
	Store_ID		  int		`json:"store_id"`
	Category_ID       int    `json:"category_id"`
	Stock_Quantity    int     `json:"stock_quantity"`
}

func New_Item(name string, price float64, category_id int, store_id int, stock_quantity int)(*Item, error) {
	return &Item{
	Name:           name,
	Price:          price,
	Category_ID:    category_id,
	Store_ID: store_id,
	Stock_Quantity: stock_quantity,
	Created_By: 1,
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
