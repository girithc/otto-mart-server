package api

import (
	"fmt"
	"net/http"
)


func (s *Server) handleCategories(res http.ResponseWriter, req *http.Request) error {
	

	if req.Method == "GET" {
		fmt.Println("Category - (GET)")
		return s.CreateCategory(res, req)
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
	fmt.Fprintf(res, "Products Endpoint Reached.")
	if req.URL.Path == "/products/" {
		
	}
	return nil
}