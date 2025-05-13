package service

import (
	"time"
	"coupon-system/internal/db"
	"coupon-system/internal/models"
)

// ValidateCoupon checks if a coupon can be applied to the cart
func ValidateCoupon(couponCode string, cartItems []models.CartItem, orderTotal float64, timestamp time.Time) (bool, string, models.Discount, error) {
	// Fetch coupon from DB
	var coupon models.Coupon
	err := db.DB.QueryRow("SELECT * FROM coupons WHERE coupon_code = $1", couponCode).Scan(&coupon.CouponCode, &coupon.ExpiryDate, &coupon.UsageType, &coupon.ApplicableMedicineIDs, &coupon.ApplicableCategories, &coupon.MinOrderValue, &coupon.ValidTimeWindow, &coupon.TermsAndConditions, &coupon.DiscountType, &coupon.DiscountValue, &coupon.MaxUsagePerUser)
	if err != nil {
		return false, "", models.Discount{}, err
	}

	// Check if coupon is expired
	if coupon.ExpiryDate.Before(timestamp) {
		return false, "coupon expired", models.Discount{}, nil
	}

	// Check if the cart meets the minimum order value
	if orderTotal < coupon.MinOrderValue {
		return false, "order total below minimum", models.Discount{}, nil
	}

	// More validation logic goes here...

	// If all checks pass, return true and the discount
	discount := models.Discount{
		ItemsDiscount: 50, // Example, to be calculated
		ChargesDiscount: 20, // Example, to be calculated
	}

	return true, "coupon applied successfully", discount, nil
}
