package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pronto-go/types"
)

func (s *Server) Handle_Create_Item(res http.ResponseWriter, req *http.Request) error {
	
	new_req := new(types.Create_Item)
	
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in Handle_Create_Item()")
		return err
	}
	
	new_item, err := types.New_Item(new_req.Name, new_req.Price, new_req.Category_ID, new_req.Store_ID, new_req.Stock_Quantity)

	if err != nil {
		return err
	}
	item, err := s.store.Create_Item(new_item); 
	
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, item)
}

func (s *Server) Handle_Get_Items(res http.ResponseWriter, req *http.Request) error {
	items, err := s.store.Get_Items()
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, items)
}

func (s *Server) Handle_Get_Items_By_(res http.ResponseWriter, req *http.Request) error {
	items, err := s.store.Get_Items()
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, items)
}