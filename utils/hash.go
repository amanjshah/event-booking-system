package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	return string(bytes), err
}

func CheckPassword(hashedPassword, password string) bool {
	// CompareHashAndPassword returns an error if the passwords do not match...
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
