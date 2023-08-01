package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"pronto-go/types"
)



func (s *Server) CreateCategory( res http.ResponseWriter, req *http.Request) error{
	
	new_req := new(types.CreateCategory)
	fmt.Println("Req: ", req )
	
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in CreateCategory()")
		return err
	}

	fmt.Println("New Request", new_req.Name, " ", new_req.ParentCategory)

	category, err := types.NewCategory(new_req.Name, new_req.ParentCategory)

	if err != nil {
		return err
	}
	if err := s.store.CreateCategory(category); err != nil {
		return err
	}

	fmt.Println("Created Category - success")

	return WriteJSON(res, http.StatusOK, category)
	
}