package tokens

import (
	"time"

	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	symmetricKey string
	paseto       *paseto.V2
}

// NewPasetoMaker creates a new PasetoMaker instance.
func NewPasetoMaker(symmetricKey string) *PasetoMaker {
	return &PasetoMaker{
		symmetricKey: symmetricKey,
		paseto:       paseto.NewV2(),
	}
}

// CreateToken generates a new PASETO token for a specific username and role with a duration.
func (maker *PasetoMaker) CreateToken(username string, role string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, role, duration, TokenTypeAccessToken)
	if err != nil {
		return "", err
	}

	token, err := maker.paseto.Encrypt([]byte(maker.symmetricKey), payload, nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

// VerifyToken checks if the PASETO token is valid and returns the payload.
func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	var payload Payload
	err := maker.paseto.Decrypt(token, []byte(maker.symmetricKey), &payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if time.Now().After(payload.ExpiredAt) {
		return nil, ErrExpiredToken
	}
	return &payload, nil
}
