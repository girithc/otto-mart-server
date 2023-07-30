package api

import (
	"net/http"
)

type Server struct {
	listen_address string
}

func NewServer(listen_address string) *Server {
	return &Server{
		listen_address: listen_address,
	}
}

func (s *Server) Run() {
	http.HandleFunc("/categories", handleCategories)
	http.HandleFunc("/products", handleProducts)

	http.ListenAndServe(s.listen_address, nil)
}




