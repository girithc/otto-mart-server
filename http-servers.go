package main

import (
	"fmt"
	"net/http"
)

func getCategoriesList(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(res, "Categories List \n")
}

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()

        for k, v := range values {
            fmt.Println(k, " => ", v)
        }
	})

    http.HandleFunc("/categories", getCategoriesList)
	http.ListenAndServe(":3000", nil)
}