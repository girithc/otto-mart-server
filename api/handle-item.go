package api

import (
	"encoding/json"
	"fmt"
	"io"
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
	
	result := new(types.Delete_Item)
	err := json.NewDecoder(req.Body).Decode(result)
	if err == io.EOF {
		items, err := s.store.Get_Items()
		if err != nil {
			return err
		}

		return WriteJSON(res, http.StatusOK, items)
	}

	if err != nil {
		return err
	}
	
	items, err := s.store.Get_Item_By_ID(result.ID)
	if err != nil {
		return err
	}
	return WriteJSON(res, http.StatusOK, items)
}

func (s *Server) Handle_Update_Item(res http.ResponseWriter, req *http.Request) error {


	new_req := new(types.Update_Item)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error Decode Item()")
		return err
	}

	item, err := s.store.Get_Item_By_ID(new_req.ID)
	if err != nil {
		return err
	}

	
	if len(new_req.Name) == 0 {
		new_req.Name = item.Name
	} 
	if new_req.Price == 0 {
		new_req.Price = item.Price
	}
	if new_req.Stock_Quantity < 0 {
		new_req.Stock_Quantity = item.Stock_Quantity
	}
  	if new_req.Category_ID == 0 {
		new_req.Category_ID = item.Category_ID
	} else {
		_, err := s.store.Get_Item_By_ID(new_req.Category_ID)
		if err != nil {
			return err
		}
	}

	updated_item, err := s.store.Update_Item(new_req)
	if err != nil {
		return err
	}


	return WriteJSON(res, http.StatusOK, updated_item)
}

func (s *Server) Handle_Delete_Item(res http.ResponseWriter, req *http.Request) error {
	
	new_req := new(types.Delete_Item)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error Decode Handle_Delete_Item()")
		return err
	}

	if err := s.store.Delete_Item(new_req.ID); err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, map[string]int{"deleted": new_req.ID})
}