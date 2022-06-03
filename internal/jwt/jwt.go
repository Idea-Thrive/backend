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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": username,
		"exp":     expirationDate,
	})

	signedToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", 0, err
	}

	return signedToken, expirationDate, nil
}
