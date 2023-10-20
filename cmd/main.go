package main

import (
	"net/http"

	"github.com/Mateusnasciment/golang-hexagonal-architecture",
)

func main() {
	controller := api.NewController()

	http.HandleFunc("/endpoint", controller.YourHandler)

	// Iniciar o servidor HTTP
	http.ListenAndServe(":8080", nil)
}
