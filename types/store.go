package types

import (
	"time"
)

type Store struct {
	ID                int       `json:"id"`
	Name				string	`json:"name"`
	Address				string	`json:"address"`
	Created_At			time.Time  `json:"created_at"`
	Created_By			int  `json:"created_by"`
}

type Create_Store struct {
	Name				string	`json:"name"`
	Address					string	`json:"address"`
}

type Update_Store struct {
	ID		int `json:"id"`
	Name		string `json:"name"`
	Address	    	string `json:"address"`
}

type Delete_Store struct {
	ID		int `json:"id"`
}

func New_Store(name string, address string)(*Store, error) {
	return &Store{
		Name: name,
		Address: address,
		Created_By: 1,
	}, nil
}

func New_Update_Store(name string, address string, id int)(*Update_Store, error) {
	return &Update_Store{
		Name: name,
		Address: address,
		ID: id,
	}, nil
}
