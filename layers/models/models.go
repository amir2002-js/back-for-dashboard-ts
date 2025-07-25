package models

import (
	"gorm.io/gorm"
	"paysee2/constants"
	"time"
)

type User struct {
	ID           uint
	Email        string
	Username     string
	//PhoneNumber  string
	Role         constants.Role
	PasswordHash string
	IsVerified   bool
	TwoFaEnabled bool
	KycStatus    constants.UserKycStatue
	LastLoginAt  time.Time
}

type KYC struct {
	gorm.Model
	UserID       uint
	DocumentType constants.KYCDocType
	FileUrl      string
	Status       constants.KYCStatuses
	UploadedAt   time.Time
	User         User
}

type TransAction struct {
	gorm.Model
	GatewayID      uint
	TxHash         string
	AmountReceived int
	Currency       int
	Network        int
	Status         int
	ConfirmedAt    time.Time
}

type LoginHistory struct {
	gorm.Model
	UserID    uint
	IPAddress string
	UserAgent string
	Location  string
	Success   bool
	Timestamp time.Time
	User      User
}
