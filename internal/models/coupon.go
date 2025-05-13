package handlers

import (
	"time"
	"fmt"
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
	Code              string      `json:"code"`
	UsageType         UsageType  `json:"usage_type"`
	MaxUsagePerUser   *int       `json:"max_usage_per_user,omitempty"` // Always optional for all types
	MinOrderValue     *float64   `json:"min_order_value,omitempty"`     // Optional, based on usage type
	TermsAndConditions string     `json:"terms_and_conditions"`
	ExpiryDate        *time.Time `json:"expiry_date,omitempty"`         // Optional, based on usage type
	Discounts         []Discount `json:"discounts"`   // Represents the related discounts
	TimeWindows       []TimeWindow `json:"time_windows,omitempty"` // Represents the related time windows for time-based coupons
}

// Discount struct for discount details
type Discount struct {
	DiscountType  string  `json:"discount_type"` // 'items' or 'charges'
	DiscountValue float64 `json:"discount_value"`
}

// TimeWindow struct for time-based coupon validity period
type TimeWindow struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

// ValidateCoupon validates the input coupon based on usage type
func (c *Coupon) Validate() error {
	// Check if usage type is valid
	switch c.UsageType {
	case OneTime, MultiUse:
		// For one_time or multi_use, max_usage_per_user is required
		if c.MaxUsagePerUser == nil {
			return fmt.Errorf("MaxUsagePerUser is required for usage type: %s", c.UsageType)
		}
	case TimeBased:
		// For time_based, expiry_date and time_windows are required
		if c.ExpiryDate == nil || len(c.TimeWindows) == 0 {
			return fmt.Errorf("ExpiryDate and TimeWindows are required for usage type: %s", c.UsageType)
		}
	default:
		return fmt.Errorf("Invalid usage type: %s", c.UsageType)
	}

	// Check if min_order_value is provided when it's necessary (for all usage types)
	if c.MinOrderValue == nil {
		return fmt.Errorf("MinOrderValue is required")
	}

	return nil
}
