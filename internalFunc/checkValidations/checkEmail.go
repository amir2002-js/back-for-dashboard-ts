package checkValidations

import (
	"errors"
	"gorm.io/gorm"
	"paysee2/layers/models"
	"regexp"
	"strings"
)

func CheckEmail(email string, db *gorm.DB) error {
	email = strings.TrimSpace(email)

	isEmail := checkEmailFormat(email)
	if !isEmail {
		return errors.New("email format error")
	}

	err := checkEmailAvailable(email, db)
	if err != nil {
		return err
	}

	return nil
}

func checkEmailAvailable(email string, db *gorm.DB) error {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}

	if err != nil {
		return err
	}

	return errors.New("email already available")
}

func checkEmailFormat(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
