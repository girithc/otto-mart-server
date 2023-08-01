package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"pronto-go/types"
)



func (s *Server) CreateCategory( res http.ResponseWriter, req *http.Request) error{
	
	new_req := new(types.CreateCategory)
	
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