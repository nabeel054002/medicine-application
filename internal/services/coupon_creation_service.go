package service

import (
	"coupon-system/internal/db"
	"coupon-system/internal/models"
	"fmt"
)

// CreateCoupon inserts a new coupon into the database
func CreateCoupon(coupon models.Coupon) error {
	// Prepare the SQL query to insert the coupon
	query := `
		INSERT INTO coupons (coupon_code, expiry_date, usage_type, applicable_medicine_ids, 
			applicable_categories, min_order_value, valid_time_window, terms_and_conditions, 
			discount_type, discount_value, max_usage_per_user)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	_, err := db.DB.Exec(query, coupon.CouponCode, coupon.ExpiryDate, coupon.UsageType, 
		coupon.ApplicableMedicineIDs, coupon.ApplicableCategories, coupon.MinOrderValue, 
		coupon.ValidTimeWindow, coupon.TermsAndConditions, coupon.DiscountType, 
		coupon.DiscountValue, coupon.MaxUsagePerUser)
	
	if err != nil {
		return fmt.Errorf("failed to insert coupon: %w", err)
	}
	return nil
}
