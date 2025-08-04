package models

import (
	"gorm.io/gorm"
	"time"
)

// Payment یک قسط یا پرداخت برای یک تراکنش خاص را ثبت می‌کند.
type Payment struct {
	gorm.Model
	// مبلغ این قسط.
	Amount float64 `json:"amount"`
	// تاریخ و زمان دقیق پرداخت. استفاده از time.Time برای گزارش‌گیری ضروری است.
	PaymentDate time.Time `json:"paymentDate"`

	CustomerId uint `json:"customerId"`
}
