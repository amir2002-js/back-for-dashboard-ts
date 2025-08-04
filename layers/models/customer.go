package models

import (
	"gorm.io/gorm"
	"paysee2/constants"
)

// Customer اطلاعات مشتریان را ذخیره می‌کند.
type Customer struct {
	gorm.Model
	FirstName       string                 `json:"firstName"`
	LastName        string                 `json:"lastName"`
	Type            constants.CustomerType `json:"type"`
	Weight          float64                `json:"weight"`
	Totality        uint                   `json:"totality"`
	PhoneNumber     string                 `json:"phoneNumber"`
	Payment         []Payment              `json:"payment"`
	Description     string                 `json:"description" gorm:"type:text"`
	UserID          uint                   `json:"userId"`
	RemainingAmount int                    `json:"remainingAmount" gorm:"-"`
}

type ClaimCustomer struct {
	FirstName   string  `json:"firstName" binding:"required"`
	LastName    string  `json:"lastName" binding:"required"`
	PhoneNumber string  `json:"phoneNumber" binding:"required"`
	Weight      float64 `json:"weight"`
	Totality    uint    `json:"totality" binding:"required"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
}
