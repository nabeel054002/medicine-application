package handlers

import (
	"encoding/json"
	"net/http"
	"coupon-system/internal/services"
	"coupon-system/internal/models"
)

// ValidateCoupon handles the coupon validation logic
func ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	var request models.ValidateCouponRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the service to validate coupon
	isValid, message, discount, err := services.ValidateCoupon(request.CouponCode, request.CartItems, request.OrderTotal, request.Timestamp)
	if err != nil {
		http.Error(w, "Error validating coupon", http.StatusInternalServerError)
		return
	}

	// Return the response
	response := struct {
		IsValid bool             `json:"is_valid"`
		Message string           `json:"message"`
		Discount models.Discount `json:"discount,omitempty"`
	}{
		IsValid: isValid,
		Message: message,
		Discount: discount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
