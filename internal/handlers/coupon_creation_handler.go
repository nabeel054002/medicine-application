package handlers

import (
	"encoding/json"
	"net/http"
	"coupon-system/internal/services"
	"coupon-system/internal/models"
)

// CreateCoupon handles the creation of new coupons
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var couponRequest models.CreateCouponRequest
	err := json.NewDecoder(r.Body).Decode(&couponRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the input
	if couponRequest.CouponCode == "" || couponRequest.ExpiryDate.IsZero() {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Prepare the coupon model from request
	coupon := models.Coupon{
		CouponCode:           couponRequest.CouponCode,
		ExpiryDate:           couponRequest.ExpiryDate,
		UsageType:            couponRequest.UsageType,
		ApplicableMedicineIDs: couponRequest.ApplicableMedicineIDs,
		ApplicableCategories: couponRequest.ApplicableCategories,
		MinOrderValue:        couponRequest.MinOrderValue,
		ValidTimeWindow:      couponRequest.ValidTimeWindow,
		TermsAndConditions:   couponRequest.TermsAndConditions,
		DiscountType:         couponRequest.DiscountType,
		DiscountValue:        couponRequest.DiscountValue,
		MaxUsagePerUser:      couponRequest.MaxUsagePerUser,
	}

	// Call the service layer to save the coupon
	err = services.CreateCoupon(coupon)
	if err != nil {
		http.Error(w, "Error creating coupon", http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Coupon created successfully"})
}
