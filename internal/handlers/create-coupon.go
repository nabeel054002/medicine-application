package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/nabeel054002/coupon-system/internal/db"
	"net/http"
	"github.com/nabeel054002/coupon-system/internal/models"
	"log"
)

// CreateCoupon handles POST /coupons/create to create a new coupon
func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.Coupon
	// Parse the incoming JSON request
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		http.Error(w, fmt.Sprintf("Error parsing request body: %v", err), http.StatusBadRequest)
		return
	}

	// Validate the coupon struct
	if err := coupon.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("Invalid coupon data: %v", err), http.StatusBadRequest)
		return
	}

	// Insert the coupon into the database
	// Assuming InsertCoupon is a function in your DB layer that inserts the coupon into the DB
	err := insertCouponIntoDB(coupon)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting coupon into the database: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with the created coupon details
	w.WriteHeader(http.StatusCreated)
	// Send the coupon back in the response, excluding sensitive information if needed
	response := map[string]interface{}{
		"message": "Coupon created successfully",
		"coupon":  coupon,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

// insertCouponIntoDB inserts the coupon into the database (this should be implemented with your DB logic)
func insertCouponIntoDB(coupon models.Coupon) error {
	// Example SQL logic (modify according to your actual DB handling)
	query := `
		INSERT INTO coupons (coupon_code, usage_type, max_usage_per_user, min_order_value, terms_and_conditions, expiry_date)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	// Here you would replace DB.Exec with your actual database execution logic
	_, err := db.DB.Exec(query, coupon.Code, coupon.UsageType, coupon.MaxUsagePerUser, coupon.MinOrderValue, coupon.TermsAndConditions, coupon.ExpiryDate)
	if err != nil {
		log.Printf("Error inserting coupon: %v", err)
		return err
	}

	// You may need to insert data into related tables (discounts, time_windows, etc.)
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

	// Insert applicable categories and medicines if provided
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

// insertDiscountIntoDB inserts a discount into the database
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

// insertApplicableMedicines inserts applicable medicines for the coupon
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
