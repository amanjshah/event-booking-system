package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 3).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(formatToken(token), func(token *jwt.Token) (any, error) {
		// someField.(/*some type*/) is special Go syntax to verify that a field is of a certain type
		// Returns 2 values: the first is the actual field (jwt.SigningMethodHS256 in this case), and the second is a boolean indicating whether the check passed.
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method! ")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return -1, errors.New("Failed to parse token. ")
	}

	if !parsedToken.Valid {
		return -1, errors.New("Invalid token. ")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("Invalid token claims. ")
	}
	// Use the type checking syntax to tell Go to store the variables as their actual type rather than of type interface{}.
	userId := int64(claims["userId"].(float64))

	return userId, nil
}

func formatToken(token string) string {
	if len(token) < 7 {
		return token
	}
	if token[:7] == "Bearer " {
		return token[7:]
	}
	return token
}
