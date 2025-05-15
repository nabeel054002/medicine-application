package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nabeel054002/coupon-system/internal/db"
	"github.com/nabeel054002/coupon-system/internal/models"
	"github.com/nabeel054002/coupon-system/internal/utils"
)

// handlers/coupons.go
func GetApplicableCoupons(w http.ResponseWriter, r *http.Request) {
	var req models.ApplicableCouponsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		http.Error(w, "Invalid timestamp format", http.StatusBadRequest)
		return
	}

	categoryPlaceholders := utils.Placeholders(len(req.CartItems))
	medicinePlaceholders := utils.Placeholders(len(req.CartItems))

	query := `
		SELECT DISTINCT c.coupon_code, d.discount_value
		FROM coupons c
		LEFT JOIN coupon_applicable_categories cc ON c.coupon_code = cc.coupon_code
		LEFT JOIN coupon_applicable_medicines cm ON c.coupon_code = cm.coupon_code
		LEFT JOIN discounts d ON c.coupon_code = d.coupon_code
		LEFT JOIN time_windows tw ON c.coupon_code = tw.coupon_code
		WHERE c.expiry_date > ?
		AND (c.min_order_value IS NULL OR c.min_order_value <= ?)
		AND (
			cc.category_id IN (
				SELECT id FROM categories WHERE name IN (` + categoryPlaceholders + `)
			)
			OR cm.medicine_id IN (` + medicinePlaceholders + `)
		)
		-- Usage type check (only apply time windows if usage_type requires it)
		AND (
			(c.usage_type != 'time_based' OR (tw.start_time IS NULL OR ? >= tw.start_time))
			AND (tw.end_time IS NULL OR ? <= tw.end_time)
		)
		AND (
			c.usage_type != 'one_time'
			OR (cu.user_id = ? AND (cu.usage_count IS NULL OR cu.usage_count = 0))
		)
	`

	args := []interface{}{timestamp, req.OrderTotal}
	for _, item := range req.CartItems {
		args = append(args, item.Category)
	}
	for _, item := range req.CartItems {
		args = append(args, item.ID)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var applicable []models.Coupon

	for rows.Next() {
		var code string
		var discountType string
		var discountValue float64

		if err := rows.Scan(&code, &discountType, &discountValue); err == nil {
			coupon := models.Coupon{
				Code: code,
				Discounts: []models.Discount{
					{
						DiscountType:  discountType,
						DiscountValue: discountValue,
					},
				},
			}
			applicable = append(applicable, coupon)
		}
	}

	json.NewEncoder(w).Encode(models.ApplicableCouponsResponse{
		ApplicableCoupons: applicable,
	})
}
