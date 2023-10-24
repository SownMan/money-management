package util

import (
	"errors"
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password : %w", err)
	}
	return string(hashPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ValidatePassword(password string) error {
	isMoreThan8 := len(password) > 8 && len(password) < 30

	var isLower, isUpper bool

	for _, r := range password {
		if !isLower && unicode.IsLower(r) {
			isLower = true
		}

		if !isUpper && unicode.IsUpper(r) {
			isUpper = true
		}
	}

	isValid := isMoreThan8 && isLower && isUpper

	if !isValid {
		return nil
	}
	return errors.New("invalid password")
}
