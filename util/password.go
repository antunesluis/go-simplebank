package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassWord return the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassWord, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password %w", err)
	}
	return string(hashedPassWord), nil
}

func CheckPassword(password string, hashedPassWord string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassWord), []byte(password))
}
