package jwt

import "time"

type Payload struct {
	Username  string    `json:"username,omitempty"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt int64     `json:"expires_at,omitempty"`
}

// Valid checks if the payload is valid.
func (p Payload) Valid() error {
	if p.IssuedAt.Before(time.Now()) && time.Unix(p.ExpiresAt, 0).Before(time.Now()) {
		return errInvalidPayload
	}

	return nil
}
