package routes

import (
	"github.com/gorilla/mux"
	"github.com/nabeel054002/coupon-system/internal/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/nabeel054002/coupon-system/docs"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	
	router.HandleFunc("/coupons/create", handlers.CreateCoupon).Methods("POST")
	// router.HandleFunc("/coupons/{code}", handlers.GetCouponByCode).Methods("GET")
	router.HandleFunc("/coupons/applicable", handlers.GetApplicableCoupons).Methods("POST")
	router.HandleFunc("/coupons/validate", handlers.ValidateCoupon).Methods("POST")

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
