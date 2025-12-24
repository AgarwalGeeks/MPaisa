package tokens

import "time"

// Maker is an interface for managing tokens.
type Maker interface {
	// CreateToken generates a new token for a specific username and role with a duration.
	CreateToken(username string, role string, duration time.Duration) (string, error)

	// VerifyToken checks if the token is valid and returns the payload.
	VerifyToken(token string) (*Payload, error)
}
