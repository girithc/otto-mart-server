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
    
    fmt.Println("Store ", store)
    

	

    server := api.NewServer(":3000")
    server.Run()
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}