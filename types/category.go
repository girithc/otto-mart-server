package types

import (
	"time"
)

type Category struct {
	ID                int       `json:"id"`
	Name              string    `json:"firstName"`
	ParentCategory    bool    `json:"parentCategory"`
	Number			  int64     `json:"number"`
	CreatedAt         time.Time `json:"createdAt"`
}

type CreateCategory struct {

	Name   string    `json:"name"`
	ParentCategory bool `json:"parentCategory"`

}

func NewCategory(name string, parentCategory bool, ) (*Category, error) {
	return &Category{
		Name: name, 
		ParentCategory: parentCategory,
		Number: 1,
		CreatedAt: time.Now().UTC(),
	}, nil
}