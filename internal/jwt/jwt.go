package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var errInvalidPayload = errors.New("invalid payload")

type JWT struct {
	Expiration time.Duration
	Secret     string
}

func NewJWT(cfg Config) *JWT {
	return &JWT{Expiration: cfg.Expiration, Secret: cfg.Secret}
}

func (j JWT) Generate(username string) (string, int64, error) {
	expirationDate := time.Now().Add(j.Expiration).Unix()

	payload := &Payload{
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiresAt: expirationDate,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", 0, err //nolint:wrapcheck
	}

	return signedToken, expirationDate, nil
}
