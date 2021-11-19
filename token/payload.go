package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Diferentes tipos de error retornados por la función VerifyToken
var (
	ErrInvalidToken = errors.New("token invalido")
	ErrExpiredToken = errors.New("token expirado")
)

// Payload contiene los datos de carga útil del token
type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewPayload crea un nuevo token payload con un nombre específico y duración.
func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

// Valid verifica si el token no es válido.
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
