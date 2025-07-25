package usecase

import (
	"paysee2/layers/models"
)

// اینجا قرارداد های برامه نوشته میشه

type UserRepository interface {
	CreateUser(user *models.User) error
	UpdateUser() error
	DeleteUser(userID int) error
	FindByUserID(id int) error
}

type LoginHistoryRepository interface {
	CreateLoginHistory() error
}

type TransActionRepository interface {
	CreateTransAction() error
}

type KYCRepository interface {
	CreateKYC() error
}
