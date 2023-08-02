package types

import (
	"time"
)

type Store struct {
	Store_ID                int       `json:"store_id"`
	Store_Name				string	`json:"store_name"`
	Address					string	`json:"address"`
	Created_At			time.Time  `json:"created_at"`
	Created_By			int  `json:"created_by"`
}

type Create_Store struct {
	Store_Name				string	`json:"store_name"`
	Address					string	`json:"address"`
}

type Update_Store struct {
	Store_ID		int `json:"store_id"`
	Store_Name		string `json:"store_name"`
	Address	    	string `json:"address"`
}

type Delete_Store struct {
	Store_ID		int `json:"store_id"`
}

func New_Store(store_name string, address string)(*Store, error) {
	return &Store{
		Store_Name: store_name,
		Address: address,
		Created_By: 1,
	}, nil
}

func New_Update_Store(store_name string, address string, store_id int)(*Update_Store, error) {
	return &Update_Store{
		Store_Name: store_name,
		Address: address,
		Store_ID: store_id,
	}, nil
}

