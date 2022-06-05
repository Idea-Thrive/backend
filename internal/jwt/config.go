package jwt

import "time"

type Config struct {
	Expiration time.Duration `koanf:"expiration"`
	Secret     string        `koanf:"secret"`
}
