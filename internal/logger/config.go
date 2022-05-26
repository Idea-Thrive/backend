package logger

// Config struct.
type Config struct {
	Level string `validate:"required" koanf:"level"`
}
