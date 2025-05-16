package models

type ValidateCouponRequest struct {
	CouponCode string  `json:"coupon_code"`
	UserID     string  `json:"user_id"`
	OrderTotal float64 `json:"order_total"`
	Timestamp  string  `json:"timestamp"`
}

type ValidateCouponResponse struct {
	Valid  bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}