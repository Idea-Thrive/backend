package mysql

// Config struct.
type Config struct {
	Host string `koanf:"host"`
	User string `koanf:"user"`
	Pass string `koanf:"pass"`
	Port string `koanf:"port"`
	Name string `koanf:"name"`
}
