package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword devuelve el hash bcrypt de la contraseña
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to has password: %w", err)
	}
	return string(hashedPassword), nil
}

// CheckPassword verifica si la contraseña porporcionada es correcta o no
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
