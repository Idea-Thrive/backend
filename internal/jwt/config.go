package jwt

import "time"

// Config struct.
type Config struct {
	Expiration time.Duration `koanf:"expiration"`
	Secret     string        `koanf:"secret"`
}
