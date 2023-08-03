package types

import "time"

type Category struct {
	ID		int `json:"id"`
	Name	string `json:"name"`
	Created_At time.Time `json:"created_at"`
	Created_By int `json:"created_by"`
}

type Create_Category struct {
	Name	string `json:"name"`
}

type Update_Category struct {
	ID		int `json:"id"`
	Name		string `json:"name"`
}

type Delete_Category struct {
	ID		int `json:"id"`
}

func New_Category(name string) (*Category, error) {
	return &Category{
	Name:       name,
	Created_By: 1,
}, nil
}

func New_Update_Category(name string, id int)(*Update_Category, error) {
	return &Update_Category{
	Name: name,
	ID:   id,
}, nil
}