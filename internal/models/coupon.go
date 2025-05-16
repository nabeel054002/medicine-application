package models

import (
	"time"
)

// UsageType represents the usage type of a coupon
type UsageType string

const (
	OneTime     UsageType = "one_time"
	MultiUse    UsageType = "multi_use"
	TimeBased   UsageType = "time_based"
)

// Coupon struct now reflects optional fields based on UsageType
type Coupon struct {
	Code                    string        `json:"code"`
	UsageType               UsageType    `json:"usage_type"`
	MaxUsagePerUser         *int         `json:"max_usage_per_user,omitempty"` // Always optional for all types
	MinOrderValue           *float64     `json:"min_order_value,omitempty"`     // Optional, based on usage type
	TermsAndConditions      string       `json:"terms_and_conditions"`
	ExpiryDate              *time.Time   `json:"expiry_date,omitempty"`         // Optional, based on usage type
	Discounts               []Discount   `json:"discounts"`                     // Represents the related discounts
	TimeWindows             []TimeWindow `json:"time_windows,omitempty"`        // Represents the related time windows for time-based coupons
	ApplicableForCategories []int        `json:"applicable_for_categories,omitempty"` // List of category IDs where the coupon applies
	ApplicableForMedicineIds []string    `json:"applicable_for_medicine_ids,omitempty"` // List of medicine IDs where the coupon applies
}

// Discount struct for discount details
type Discount struct {
	DiscountType  string  `json:"discount_type"` // 'items' or 'charges' // can make this also into a separate type but tbh its fine for now
	DiscountValue float64 `json:"discount_value"`
}

// TimeWindow struct for time-based coupon validity period
type TimeWindow struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

