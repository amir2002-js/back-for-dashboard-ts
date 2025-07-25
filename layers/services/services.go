package services

import (
	"paysee2/layers/usecase"
)

type UserService struct {
	repo usecase.UserRepository
}

type KYCService struct {
	repo usecase.KYCRepository
}

type TransActionService struct {
	repo usecase.TransActionRepository
}

type LoginHistoryService struct {
	repo usecase.LoginHistoryRepository
}
