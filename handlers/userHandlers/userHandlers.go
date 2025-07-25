package userHandlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"paysee2/constants"
	"paysee2/internalFunc/checkValidations"
	"paysee2/internalFunc/password"
	"paysee2/layers/models"
	"paysee2/layers/services"
	"strings"
)

type UserHandlers struct {
	userServ services.UserService
	db       *gorm.DB
}

func NewUserHandlers(userServ services.UserService, db *gorm.DB) *UserHandlers {
	return &UserHandlers{userServ: userServ, db: db}
}

func (handler *UserHandlers) CreateUser(c *gin.Context) {
	claim := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Username string `json:"username" binding:"required"`
	}{}

	err := c.ShouldBind(&claim)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": constants.ErrorBindUser})
		return
	}

	err = checkValidations.CheckEmail(claim.Email, handler.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := password.HashPassword(claim.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mainUser := &models.User{
		Email:        strings.TrimSpace(claim.Email),
		Username:     strings.TrimSpace(claim.Username),
		PasswordHash: hashedPassword,
	}

	err = handler.userServ.CreateUser(mainUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": constants.ErrorCreateUser})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"result": constants.SuccessFull, "accessToken": "", "refreshToken": ""})
}
