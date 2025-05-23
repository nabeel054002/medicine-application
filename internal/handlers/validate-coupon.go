package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/nabeel054002/coupon-system/internal/db"
	"github.com/nabeel054002/coupon-system/internal/models"
)

// @Summary Validate a coupon
// @Description Validates a coupon code based on order total, timestamp, and user usage
// @Tags coupons
// @Accept json
// @Produce json
// @Param request body models.ValidateCouponRequest true "Coupon Validation Request"
// @Success 200 {object} models.ValidateCouponResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /coupons/validate [post]
func ValidateCoupon(w http.ResponseWriter, r *http.Request) {
	var req models.ValidateCouponRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		http.Error(w, "Invalid timestamp format", http.StatusBadRequest)
		return
	}

	var (
		expiryDate       time.Time
		minOrderValue    *float64
		usageType        string
		maxUsagePerUser  *int
	)

	err = db.DB.QueryRow(`
		SELECT expiry_date, min_order_value, usage_type, max_usage_per_user
		FROM coupons
		WHERE coupon_code = ?
	`, req.CouponCode).Scan(&expiryDate, &minOrderValue, &usageType, &maxUsagePerUser)
	if err != nil {
		http.Error(w, "Coupon not found", http.StatusNotFound)
		return
	}

	if timestamp.After(expiryDate) {
		json.NewEncoder(w).Encode(models.ValidateCouponResponse{
			Valid:  false,
			Reason: "Coupon expired",
		})
		return
	}

	print("usageType: ", usageType)

	if minOrderValue != nil && req.OrderTotal < *minOrderValue {
		json.NewEncoder(w).Encode(models.ValidateCouponResponse{
			Valid:  false,
			Reason: "Order total less than minimum required for coupon",
		})
		return
	}

	if usageType == "one_time" || usageType == "multi_use" {
		var usageCount int
	err = db.DB.QueryRow(`
		SELECT usage_count FROM coupon_usages
		WHERE coupon_code = ? AND user_id = ?
	`, req.CouponCode, req.UserID).Scan(&usageCount)

	if err == sql.ErrNoRows {
		usageCount = 0 // first-time user
	} else if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

		if usageType == "one_time" && usageCount > 0 {
			json.NewEncoder(w).Encode(models.ValidateCouponResponse{
				Valid:  false,
				Reason: "Coupon already used by this user",
			})
			return
		}

		if usageType == "multi_use" && maxUsagePerUser != nil && usageCount >= *maxUsagePerUser {
			json.NewEncoder(w).Encode(models.ValidateCouponResponse{
				Valid:  false,
				Reason: "Coupon usage limit reached for this user",
			})
			return
		}
	}

	// Step 3: If valid, increment usage count (upsert pattern)
	_, err = db.DB.Exec(`
		INSERT INTO coupon_usages (coupon_code, user_id, usage_count)
		VALUES (?, ?, 1)
		ON CONFLICT(coupon_code, user_id) DO UPDATE SET usage_count = usage_count + 1
	`, req.CouponCode, req.UserID)

	if err != nil {
		http.Error(w, "Failed to update coupon usage", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.ValidateCouponResponse{
		Valid: true,
	})
}
