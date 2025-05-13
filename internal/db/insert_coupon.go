package db

import (
	"github.com/nabeel054002/coupon-system/internal/models"
	"fmt"
	"log"
)

// InsertCoupon inserts a new coupon into the database
func InsertCoupon(coupon models.Coupon) error {
	query := `
		INSERT INTO coupons (code, discount_type, discount_value, valid_from, valid_until)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := DB.Exec(query, coupon.Code, coupon.DiscountType, coupon.DiscountValue, coupon.ValidFrom, coupon.ValidUntil)
	if err != nil {
		log.Printf("Error inserting coupon: %v", err)
		return fmt.Errorf("unable to insert coupon: %w", err)
	}
	return nil
}
