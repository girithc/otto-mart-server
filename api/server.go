package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pronto-go/storage"
)

type Server struct {
	listen_address string
	store *storage.PostgresStore
}

func NewServer(listen_address string, store *storage.PostgresStore) *Server {
	return &Server{
		listen_address: listen_address,
		store: store,
	}
}

func (s *Server) Run() {
	

	http.HandleFunc("/store/available", makeHTTPHandleFunc(s.handleStoreClient))
	http.HandleFunc("/store", makeHTTPHandleFunc(s.handleStoreManager))

	http.HandleFunc("/higher-level-category", makeHTTPHandleFunc(s.handleHigherLevelCategory))
	
	http.HandleFunc("/category-higher-level-mapping", makeHTTPHandleFunc(s.handleCategoryHigherLevelMapping))

	http.HandleFunc("/category", makeHTTPHandleFunc(s.handleCategory))
	http.HandleFunc("/item", makeHTTPHandleFunc(s.handleItem))

	fmt.Println("Listening PORT", s.listen_address)

	http.ListenAndServe(s.listen_address, nil)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
} 




