package handlers

import (
	"github.com/nabeel054002/coupon-system/internal/db" // Assuming this is where the DB logic resides
	"github.com/nabeel054002/coupon-system/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)


// CreateCoupon handles POST /coupons/create
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.Coupon

	// Parse the incoming JSON request body into the Coupon struct
	err := json.NewDecoder(r.Body).Decode(&coupon)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid input: %v", err), http.StatusBadRequest)
		return
	}

	// Validate the input data
	if coupon.Code == "" || coupon.DiscountType == "" || coupon.DiscountValue <= 0 {
		http.Error(w, "Missing or invalid required fields", http.StatusBadRequest)
		return
	}

	// Insert the coupon into the database
	err = db.InsertCoupon(coupon)
	if err != nil {
		log.Printf("Error inserting coupon into DB: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon created successfully"})
}
