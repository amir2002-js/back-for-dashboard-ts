package models

import (
	"gorm.io/gorm"
	"paysee2/constants"
)

type User struct {
	gorm.Model
	Email        string         `json:"email"`
	Username     string         `json:"username"`
	Role         constants.Role `json:"role"`
	PasswordHash string         `json:"passwordHash"`
}

type Creditor struct {
	gorm.Model
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	PhoneNumber      string `json:"phoneNumber"`
	TotalMoney       int    `json:"totalMoney"`
	CreditCardNumber string `json:"creditCard"`
	UserID           uint   `json:"userId"`
	User             User
}
type PaymentCreditor struct {
	gorm.Model
	CreditorID uint `json:"creditorID"`
	Pay        uint `json:"pay"`
	Creditor   Creditor
}

type Debtor struct {
	gorm.Model
	TotalMoney  int    `json:"totalMoney"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	UserID      uint   `json:"userId"`
	User        User
}

type PaymentDebtor struct {
	gorm.Model
	DebtorID uint `json:"debtorID"`
	Pay      uint `json:"pay"`
	Debtor   Debtor
}
