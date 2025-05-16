package models

type ApplicableCouponsRequest struct {
	CartItems  []CartItem `json:"cart_items"`
	OrderTotal float64    `json:"order_total"`
	Timestamp  string     `json:"timestamp"`
	UserID    *string     `json:"user_id,omitempty"` // optional for now
}

type CartItem struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

type ApplicableCouponsResponse struct {
	ApplicableCoupons []Coupon `json:"applicable_coupons"`
}
