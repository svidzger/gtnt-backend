package main

import (
	"log"
	"net/http"

	"github.com/svidzger/gtnt-backend/db"
	"github.com/svidzger/gtnt-backend/handlers"
)

func main() {
	// Connect to the database
	db.ConnectDB()

	// Initialize a new ServeMux
	mux := http.NewServeMux()

	// Define authentication-related API routes
	mux.HandleFunc("/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/auth/login", handlers.LoginHandler)

	// Start the server
	log.Println("Starting server...")
	log.Println("Server is running at http://localhost:8080/")
	log.Println("Press Ctrl + C to stop the server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
