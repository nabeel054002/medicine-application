package models

// ApplicableCouponsRequest represents the input to the GET /coupons/applicable endpoint
type ApplicableCouponsRequest struct {
	CartItems  []CartItem `json:"cart_items"`
	OrderTotal float64    `json:"order_total"`
	Timestamp  string     `json:"timestamp"`
}

// CartItem represents an item in the user's cart
type CartItem struct {
	ID       string `json:"id"`
	Category string `json:"category"`
}

// ApplicableCouponsResponse is the output structure containing matched coupons
type ApplicableCouponsResponse struct {
	ApplicableCoupons []Coupon `json:"applicable_coupons"`
}
