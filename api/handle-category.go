package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pronto-go/types"
)



func (s *Server) CreateCategory( res http.ResponseWriter, req *http.Request) error{
	
	new_req := new(types.Create_Category)
	
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in CreateCategory()")
		return err
	}
	
	category, err := types.NewCategory(new_req.Name, new_req.ParentCategory)

	if err != nil {
		return err
	}
	if err := s.store.CreateCategory(category); err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, category)
}

func (s *Server) GetCategory( res http.ResponseWriter, req *http.Request) error {
	
	categories, err := s.store.GetCategory()
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, categories)
}