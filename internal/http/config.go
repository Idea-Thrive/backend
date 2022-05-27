package http

// Config struct.
type Config struct {
	Port   int    `koanf:"port"`
	Secret string `koanf:"secret"`
}
