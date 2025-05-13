package main

import (
	"log"
	"net/http"
	"github.com/nabeel054002/coupon-system/internal/routes"
	"github.com/nabeel054002/coupon-system/internal/db"
)

func main() {
	// Initialize the database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize the database: %v", err)
	}

	// Create the schema (tables)
	err = db.CreateSchema()
	if err != nil {
		log.Fatalf("Could not create the database schema: %v", err)
	}

	// Use centralized router
	router := routes.NewRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
