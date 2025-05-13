package main

import (
	"log"
	"net/http"
	"coupon-system/internal/db"
	"coupon-system/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize the database: %v", err)
	}

	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/coupons/create", handlers.CreateCoupon).Methods("POST")

	// Start the server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
