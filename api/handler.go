package api

import (
	"fmt"
	"net/http"
)

// Store

func (s *Server) handleStoreClient(res http.ResponseWriter, req *http.Request) error {

	if req.Method == "GET" {
		pincode := req.URL.Query().Get("pincode")
		fmt.Println("We deliver to Area =>", pincode)
	}

	return nil
}

func (s *Server) handleStoreManager(res http.ResponseWriter, req *http.Request) error {

	if req.Method == "GET" {
		fmt.Println("[GET] - Store(s)")
		return s.Handle_Get_Stores(res, req)
	} else if req.Method == "POST" {
		fmt.Println("[POST] - Store")
		return s.Handle_Create_Store(res, req)
	} else if req.Method == "PUT" {
		fmt.Println("[PUT] - Store")
		return s.Handle_Update_Store(res, req)
	} else if req.Method == "DELETE" {
		fmt.Println("[DELETE] - Store")
		return s.Handle_Delete_Store(res, req)
	}

	return fmt.Errorf("No Matching Path")
}


// Category

func (s *Server) handleCategory(res http.ResponseWriter, req *http.Request) error {
	

	if req.Method == "GET" {
		fmt.Println("Category - (GET)")
		return s.GetCategory(res, req)
	} else if req.Method == "POST" {
		fmt.Println("Category - (POST)")
		return s.CreateCategory(res, req)
	} else if req.Method == "PUT" {
		return nil
	} else if req.Method == "DELETE" {
		return nil
	} 

	return nil
}

// Item

func (s *Server) handleItem(res http.ResponseWriter, req *http.Request) error{
	
	return nil
}