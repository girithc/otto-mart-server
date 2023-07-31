package api

import (
	"fmt"
	"net/http"
	"pronto-go/storage"
)

type Server struct {
	listen_address string
	store storage.PostgresStore
}

func NewServer(listen_address string, store storage.PostgresStore) *Server {
	return &Server{
		listen_address: listen_address,
		store: store,
	}
}

func (s *Server) Run() {

	http.HandleFunc("/categories", s.handleCategories)
	http.HandleFunc("/products", s.handleProducts)

	fmt.Println("Server Store: ",s.store)

	http.ListenAndServe(s.listen_address, nil)
}




