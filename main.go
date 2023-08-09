package main

import (
	"fmt"
	"log"
	"pronto-go/api"
	"pronto-go/storage"
)
 

func main() {
    
    fmt.Printf("Starting Pronto-DB\n\n")

    store, err := storage.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	} 

    if err := store.Init(); err != nil {
		log.Fatal(err)
	} else {
        fmt.Println("Store.Init() is successful.")
    }

    server := api.NewServer(":3000", store)
    server.Run()
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}