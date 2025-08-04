package handlers

import (
	"gorm.io/gorm"
	"paysee2/handlers/customerHandler"
	"paysee2/handlers/userHandlers"
)

type Handlers struct {
	CustomerHandler *customerHandler.CustomerHandlers
	UserHandler     *userHandlers.UserHandlers
	DB              *gorm.DB
}

func NewHandlers(userHandler *userHandlers.UserHandlers, customHandler *customerHandler.CustomerHandlers, db *gorm.DB) *Handlers {
	return &Handlers{UserHandler: userHandler, CustomerHandler: customHandler, DB: db}
}
