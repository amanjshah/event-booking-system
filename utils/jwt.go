package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userId,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
