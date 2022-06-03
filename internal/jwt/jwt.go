package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Expiration time.Duration
	Secret     string
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
