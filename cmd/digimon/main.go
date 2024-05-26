package main

import (
	"fmt"
	"net/http"

	"github.com/urechd/digimon/internal/routes"
)

func main() {
	fmt.Println("Hello DigiDestined")

	router := routes.NewRouter()

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
