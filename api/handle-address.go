package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/girithc/pronto-go/types"
)

func (s *Server) Handle_Create_Address(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.Create_Address)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in Handle_Create_Address()")
		return err
	}

	addr, err := s.store.Create_Address(new_req)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, addr)
}

func (s *Server) Handle_Get_Address_By_Customer_Id(res http.ResponseWriter, req *http.Request) error {
	new_req := new(types.Address_Customer_Id)
	if err := json.NewDecoder(req.Body).Decode(new_req); err != nil {
		fmt.Println("Error in Decoding req.body in Handle_Get_Address_By_Customer_Id()")
		return err
	}

	addrs, err := s.store.Get_Addresses_By_Customer_Id(new_req.Customer_Id)
	if err != nil {
		return err
	}

	return WriteJSON(res, http.StatusOK, addrs)
}