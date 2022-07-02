package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var errInvalidPayload = errors.New("invalid payload")

// JWT struct.
type JWT struct {
	Expiration time.Duration
	Secret     string
}

// NewJWT function.
func NewJWT(cfg Config) *JWT {
	return &JWT{Expiration: cfg.Expiration, Secret: cfg.Secret}
}

// Generate function.
func (j JWT) Generate(email string) (string, int64, error) {
	expirationDate := time.Now().Add(j.Expiration).Unix()

	payload := &Payload{
		Email:     email,
		IssuedAt:  time.Now(),
		ExpiresAt: expirationDate,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expirationDate, nil
}

// Verify function.
func (j JWT) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(j.Secret), nil
	}

	payload := new(Payload)
	if _, err := jwt.ParseWithClaims(token, payload, keyFunc); err != nil {
		return nil, err
	}

	return payload, nil
}
