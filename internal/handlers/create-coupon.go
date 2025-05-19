package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"

	"github.com/nabeel054002/coupon-system/internal/db"
	"github.com/nabeel054002/coupon-system/internal/models"
)

// CreateCoupon godoc
// @Summary Create a new coupon
// @Description Adds a new coupon with applicable discounts, time windows, and restrictions.
// @Tags coupons
// @Accept json
// @Produce json
// @Param coupon body models.Coupon true "Coupon to create"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /coupons/create [post]
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.Coupon
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := coupon.IsCoupon(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid coupon data: %v", err), http.StatusBadRequest)
		return
	}

	err := insertCouponIntoDB(coupon)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting coupon into the database: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Coupon created successfully",
		"coupon":  coupon,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

func insertCouponIntoDB(coupon models.Coupon) error {
	query := `
		INSERT INTO coupons (coupon_code, usage_type, max_usage_per_user, min_order_value, terms_and_conditions, expiry_date)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := db.DB.Exec(query, coupon.Code, coupon.UsageType, coupon.MaxUsagePerUser, coupon.MinOrderValue, coupon.TermsAndConditions, coupon.ExpiryDate)
	if err != nil {
		log.Printf("Error inserting coupon: %v", err)
		return err
	}

	for _, discount := range coupon.Discounts {
		err := insertDiscountIntoDB(coupon.Code, discount)
		if err != nil {
			return fmt.Errorf("Error inserting discount: %v", err)
		}
	}

	for _, timeWindow := range coupon.TimeWindows {
		err := insertTimeWindowIntoDB(coupon.Code, timeWindow)
		if err != nil {
			return fmt.Errorf("Error inserting time window: %v", err)
		}
	}

	err = insertApplicableCategories(coupon.Code, coupon.ApplicableForCategories)
	if err != nil {
		return fmt.Errorf("Error inserting applicable categories: %v", err)
	}

	err = insertApplicableMedicines(coupon.Code, coupon.ApplicableForMedicineIds)
	if err != nil {
		return fmt.Errorf("Error inserting applicable medicines: %v", err)
	}

	return nil
}

func insertDiscountIntoDB(couponCode string, discount models.Discount) error {
	query := `
		INSERT INTO discounts (coupon_code, discount_type, discount_value)
		VALUES (?, ?, ?)
	`
	_, err := db.DB.Exec(query, couponCode, discount.DiscountType, discount.DiscountValue)
	return err
}

// insertTimeWindowIntoDB inserts a time window into the database for time-based coupons
func insertTimeWindowIntoDB(couponCode string, timeWindow models.TimeWindow) error {
	query := `
		INSERT INTO time_windows (coupon_code, start_time, end_time)
		VALUES (?, ?, ?)
	`
	_, err := db.DB.Exec(query, couponCode, timeWindow.StartTime, timeWindow.EndTime)
	return err
}

// insertApplicableCategories inserts applicable categories for the coupon
func insertApplicableCategories(couponCode string, categoryIds []int) error {
	for _, categoryId := range categoryIds {
		query := `
			INSERT INTO coupon_applicable_categories (coupon_code, category_id)
			VALUES (?, ?)
		`
		_, err := db.DB.Exec(query, couponCode, categoryId)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertApplicableMedicines(couponCode string, medicineIds []string) error {
	for _, medicineId := range medicineIds {
		query := `
			INSERT INTO coupon_applicable_medicines (coupon_code, medicine_id)
			VALUES (?, ?)
		`
		_, err := db.DB.Exec(query, couponCode, medicineId)
		if err != nil {
			return err
		}
	}
	return nil
}
