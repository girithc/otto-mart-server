package api

import (
	"fmt"
	"net/http"
)


func handleCategories(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Categories Endpoint Reached.")
}

func handleProducts(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Products Endpoint Reached.")
}