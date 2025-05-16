package routes

import (
	"github.com/gorilla/mux"
	"github.com/nabeel054002/coupon-system/internal/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/coupons/create", handlers.CreateCoupon).Methods("POST")
	router.HandleFunc("/coupons/{code}", handlers.GetCouponByCode).Methods("GET")
	router.HandleFunc("/coupons/applicable", handlers.GetApplicableCoupons).Methods("POST")
	router.HandleFunc("/coupons/validate", handlers.ValidateCoupon).Methods("POST")

	return router
}
