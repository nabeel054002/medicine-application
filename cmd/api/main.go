
// @title Coupon Service API
// @version 1.0
// @description API to apply and validate coupons for medicine cart
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"net/http"
	"github.com/nabeel054002/coupon-system/internal/routes"
	"github.com/nabeel054002/coupon-system/internal/db"

)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize the database: %v", err)
	}

	err = db.CreateSchema()
	if err != nil {
		log.Fatalf("Could not create the database schema: %v", err)
	}

	err = db.SeedCategories()
	if err != nil {
		log.Fatalf("Error seeding categories: %v", err)
	}

	err = db.SeedMedicines()
	if err != nil {
		log.Fatalf("Error in seeding medicines: %v", err)
	}

	// Use centralized router
	router := routes.NewRouter()

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
