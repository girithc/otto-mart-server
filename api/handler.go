package api

import (
	"fmt"
	"net/http"
)


func (s *Server) handleCategories(res http.ResponseWriter, req *http.Request) {
	

	if req.Method == "GET" {
		return 
	} else if req.Method == "POST" {
		return 
	} else if req.Method == "PUT" {
		return 
	} else if req.Method == "DELETE" {
		return 
	} else {
		return 
	}

	return 
}

func (s *Server) handleProducts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Products Endpoint Reached.")
	if req.URL.Path == "/products/" {
		if req.Method == http.MethodPost {

		} 
	}
}