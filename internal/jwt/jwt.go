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

func (j JWT) Verify(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(j.Secret), nil
	}

	payload := new(Payload)
	if _, err := jwt.ParseWithClaims(token, payload, keyFunc); err != nil {
		return nil, err //nolint:wrapcheck
	}

	return payload, nil
}
