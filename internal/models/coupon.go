package models

import "time"

// Coupon represents a coupon in the system
type Coupon struct {
	CouponCode           string    `json:"coupon_code"`
	ExpiryDate           time.Time `json:"expiry_date"`
	UsageType            string    `json:"usage_type"`
	ApplicableMedicineIDs []string  `json:"applicable_medicine_ids"`
	ApplicableCategories []string  `json:"applicable_categories"`
	MinOrderValue        float64   `json:"min_order_value"`
	ValidTimeWindow      []string  `json:"valid_time_window"`
	TermsAndConditions   string    `json:"terms_and_conditions"`
	DiscountType         string    `json:"discount_type"`
	DiscountValue        float64   `json:"discount_value"`
	MaxUsagePerUser      int       `json:"max_usage_per_user"`
}

type CartItem struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
}

type Discount struct {
	ItemsDiscount  float64 `json:"items_discount"`
	ChargesDiscount float64 `json:"charges_discount"`
}

type ValidateCouponRequest struct {
	CouponCode string      `json:"coupon_code"`
	CartItems  []CartItem  `json:"cart_items"`
	OrderTotal float64     `json:"order_total"`
	Timestamp  time.Time   `json:"timestamp"`
}


// CreateCouponRequest represents the data required to create a coupon
type CreateCouponRequest struct {
	CouponCode           string    `json:"coupon_code"`
	ExpiryDate           time.Time `json:"expiry_date"`
	UsageType            string    `json:"usage_type"`
	ApplicableMedicineIDs []string  `json:"applicable_medicine_ids"`
	ApplicableCategories []string  `json:"applicable_categories"`
	MinOrderValue        float64   `json:"min_order_value"`
	ValidTimeWindow      []string  `json:"valid_time_window"`
	TermsAndConditions   string    `json:"terms_and_conditions"`
	DiscountType         string    `json:"discount_type"`
	DiscountValue        float64   `json:"discount_value"`
	MaxUsagePerUser      int       `json:"max_usage_per_user"`
}
