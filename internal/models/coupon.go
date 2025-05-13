package models

import "time"

// Coupon represents a coupon in the system
type Coupon struct {
	Code         string    `json:"code"`
	DiscountType string    `json:"discount_type"`
	DiscountValue float64  `json:"discount_value"`
	ValidFrom    time.Time `json:"valid_from"`
	ValidUntil   time.Time `json:"valid_until"`
}