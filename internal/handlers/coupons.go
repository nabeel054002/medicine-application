package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// GetCouponByCode handles GET /coupons/{code}
func GetCouponByCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]

	json.NewEncoder(w).Encode(map[string]string{
		"message":     "GetCouponByCode placeholder",
		"coupon_code": code,
	})
}

// GetApplicableCoupons handles POST /coupons/applicable
func GetApplicableCoupons(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "GetApplicableCoupons placeholder"})
}

// ValidateCoupon handles POST /coupons/validate
func ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ValidateCoupon placeholder"})
}
