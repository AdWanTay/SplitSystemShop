package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Driver   string `envconfig:"DB_DRIVER" default:"sqlite"` // sqlite / postgres
	Host     string `envconfig:"DB_HOST"`
	Port     string `envconfig:"DB_PORT"`
	User     string `envconfig:"DB_USER"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
}

type Config struct {
	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`
	Port      string `envconfig:"PORT" default:"8080"`
	Database  DatabaseConfig
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
