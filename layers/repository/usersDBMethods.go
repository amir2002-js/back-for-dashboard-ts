package repository

import (
	"gorm.io/gorm"
	"paysee2/layers/models"
)

func NewGormDB(db *gorm.DB) *GormDB {
	return &GormDB{DB: db}
}

func (database *GormDB) CreateUser(user *models.User) error {

	return nil
}

func (database *GormDB) UpdateUser() error {
	return nil
}

func (database *GormDB) DeleteUser(userID int) error {
	return nil
}

func (database *GormDB) FindByUserID(id int) error {
	return nil
}
