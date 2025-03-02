package config

type HTTP struct {
	Host string `env:"HOST" envDefault:"0.0.0.0"`
	Port string `env:"PORT" envDefault:"8080"`
}
type DB struct {
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	Database string `env:"DATABASE"`
	Port     int    `env:"PORT"`
}

type Cfg struct {
	Env  string `env:"ENV" envDefault:".env"`
	HTTP HTTP   `envPrefix:"HTTP_"`
	DB   DB     `envPrefix:"DB_"`
}
