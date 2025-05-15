package models

// Request structure for validating a coupon
type ValidateCouponRequest struct {
	CouponCode string  `json:"coupon_code"`
	UserID     string  `json:"user_id"`
	OrderTotal float64 `json:"order_total"`
	Timestamp  string  `json:"timestamp"` // RFC3339 format
}

// Response structure for validation result
type ValidateCouponResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}