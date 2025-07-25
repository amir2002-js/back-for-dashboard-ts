package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"paysee2/constants"
	"paysee2/layers/models"
	"time"
)

func GenerateJWTRefreshTkn(user *models.User) (strToken string, err error) {
	secretRefreshTkn := os.Getenv(constants.RefreshTknEnv)
	claims := &jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err = token.SignedString([]byte(secretRefreshTkn))
	return
}
