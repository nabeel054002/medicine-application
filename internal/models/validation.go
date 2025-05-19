package models

import (
	"fmt"
	"time"
)

func validateGeneralCouponFields(c *Coupon) error {
	now := time.Now().UTC()

	if c.Code == "" {
		return fmt.Errorf("code is required")
	}
	if c.MinOrderValue == nil {
		return fmt.Errorf("min_order_value is required")
	}
	if len(c.Discounts) == 0 {
		return fmt.Errorf("at least one discount must be provided")
	}
	for i, d := range c.Discounts {
		if d.DiscountType != "inventory" && d.DiscountType != "charges" {
			return fmt.Errorf("discounts[%d]: invalid discount_type '%s'", i, d.DiscountType)
		}
		if d.DiscountValue <= 0 {
			return fmt.Errorf("discounts[%d]: discount_value must be greater than 0", i)
		}
	}
	if c.ExpiryDate != nil && c.ExpiryDate.Before(now) {
		return fmt.Errorf("cannot create a coupon that is already expired")
	}
	return nil
}

func validateOneTimeCoupon(c *Coupon) error {
	if c.MaxUsagePerUser == nil {
		return fmt.Errorf("max_usage_per_user is required for one_time usage")
	}
	if c.ExpiryDate != nil {
		return fmt.Errorf("expiry_date should not be set for one_time usage")
	}
	if len(c.TimeWindows) > 0 {
		return fmt.Errorf("time_windows should not be set for one_time usage")
	}
	return nil
}

func validateTimeBasedCoupon(c *Coupon) error {
	if len(c.TimeWindows) == 0 {
		return fmt.Errorf("at least one time_window is required for time_based usage")
	}
	for i, tw := range c.TimeWindows {
		if tw.EndTime.Before(tw.StartTime) {
			return fmt.Errorf("time_window[%d]: end_time must be after start_time", i)
		}
	}
	if c.MaxUsagePerUser != nil {
		return fmt.Errorf("max_usage_per_user is required for multi_use usage")
	}
	return nil
}

func validateMultiUseCoupon(c *Coupon) error {
	if c.MaxUsagePerUser == nil {
		return fmt.Errorf("max_usage_per_user is required for multi_use usage")
	}
	if len(c.TimeWindows) > 0 {
		return fmt.Errorf("time_windows should not be set for multi_use usage")
	}
	return nil
}


// ValidateCoupon validates the input coupon based on usage type
func (c *Coupon) IsCoupon() error {
	if err := validateGeneralCouponFields(c); err != nil {
		return err
	}

	switch c.UsageType {
	case OneTime:
		return validateOneTimeCoupon(c)
	case MultiUse:
		return validateMultiUseCoupon(c)
	case TimeBased:
		return validateTimeBasedCoupon(c)
	default:
		return fmt.Errorf("invalid usage type: %s", c.UsageType)
	}
}