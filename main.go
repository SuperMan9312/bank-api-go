package main

import (
	"log"
	"net/http"

	"bank-api/router"
)

func main() {
	// Initialize routes
	r := router.InitializeRoutes()

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
