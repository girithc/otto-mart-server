package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello World\n")
}

func categories(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "List of Product Categories")
	fmt.Fprintf(res, "name: "+req.FormValue("name"))
}

func queryParamDisplayHandler(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "name: "+req.FormValue("name"))
    io.WriteString(res, "\nphone: "+req.FormValue("phone"))
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
	http.HandleFunc("/categories", categories)
	http.HandleFunc("/example", func(res http.ResponseWriter, req *http.Request) {
        queryParamDisplayHandler(res, req)
    })
    http.ListenAndServe(":8080", nil)
}