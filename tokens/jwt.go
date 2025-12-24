package tokens

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// JWTMaker is responsible for creating and verifying JWT tokens
type JWTMaker struct {
	secretKey string
}

// NewJWTMaker creates a new JWTMaker instance.
func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{
		secretKey: secretKey,
	}
}

// CreateToken generates a new JWT token for a specific username and role with a duration.
func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, role, duration, TokenTypeAccessToken)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":         payload.ID.String(),
		"token_type": payload.Type,
		"username":   payload.Username,
		"role":       payload.Role,
		"issued_at":  payload.IssuedAt.Unix(),
		"expired_at": payload.ExpiredAt.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the JWT token is valid and returns the payload.
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(maker.secretKey), nil
	})
	if err != nil {
		return nil, ErrInvalidToken
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, ErrInvalidToken
		}
		return &Payload{
			ID:        id,
			Type:      TokenType(claims["token_type"].(float64)),
			Username:  claims["username"].(string),
			Role:      claims["role"].(string),
			IssuedAt:  time.Unix(int64(claims["issued_at"].(float64)), 0),
			ExpiredAt: time.Unix(int64(claims["expired_at"].(float64)), 0),
		}, nil
	}
	return nil, ErrInvalidToken
}
