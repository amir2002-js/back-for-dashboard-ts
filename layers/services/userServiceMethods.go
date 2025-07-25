package services

import (
	"paysee2/layers/models"
	"paysee2/layers/usecase"
)

func NewUserService(ucRepo usecase.UserRepository) *UserService {
	return &UserService{repo: ucRepo}
}

func (uService UserService) CreateUser(user *models.User) error {
	return uService.repo.CreateUser(user)
}

func (uService UserService) UpdateUser() error {
	return uService.repo.UpdateUser()
}

func (uService UserService) DeleteUser(userID int) error {
	return uService.repo.DeleteUser(userID)
}

func (uService UserService) FindByUserID(id int) error {
	return uService.repo.FindByUserID(id)
}
