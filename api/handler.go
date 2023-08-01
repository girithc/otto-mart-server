package api

import (
	"fmt"
	"net/http"
)

func (s *Server) handleStoreClient(res http.ResponseWriter, req *http.Request) error {

	if req.Method == "GET" {
		pincode := req.URL.Query().Get("pincode")
		fmt.Println("We deliver to Area =>", pincode)
	}

	return nil
}

func (s *Server) handleStoreManager(res http.ResponseWriter, req *http.Request) error {

	if req.Method == "GET" {
		storeId := req.URL.Query().Get("id")
		fmt.Println("Store Id =>", storeId)
	}

	return nil
}


func (s *Server) handleCategories(res http.ResponseWriter, req *http.Request) error {
	

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

func (s *Server) handleProducts(res http.ResponseWriter, req *http.Request) error{
	
	return nil
}