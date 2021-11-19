package token

import "time"

// Maker interfaz que administra la creación y verificación de tokens
type Maker interface {
	// CreateToken crea un nuevo token para un usuario específico y duración
	CreateToken(username string, duration time.Duration) (string, error)

	// Verifica si el token es inválido o no.
	VerifyToken(token string) (*Payload, error)
}
