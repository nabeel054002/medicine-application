package models

import (
	"time"
)

type UsageType string

const (
	OneTime     UsageType = "one_time"
	MultiUse    UsageType = "multi_use"
	TimeBased   UsageType = "time_based"
)

type Coupon struct {
	Code                    string        `json:"code"`
	UsageType               UsageType    `json:"usage_type"`
	MaxUsagePerUser         *int         `json:"max_usage_per_user,omitempty"` 
	MinOrderValue           *float64     `json:"min_order_value,omitempty"`     
	TermsAndConditions      string       `json:"terms_and_conditions"`
	ExpiryDate              *time.Time   `json:"expiry_date,omitempty"`         
	Discounts               []Discount   `json:"discounts"`                     
	TimeWindows             []TimeWindow `json:"time_windows,omitempty"`        
	ApplicableForCategories []int        `json:"applicable_for_categories,omitempty"` 
	ApplicableForMedicineIds []string    `json:"applicable_for_medicine_ids,omitempty"` 
}

type Discount struct {
	DiscountType  string  `json:"discount_type"` // 'items' or 'charges' // can make this also into a separate type but tbh its fine for now
	DiscountValue float64 `json:"discount_value"`
}

type TimeWindow struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

