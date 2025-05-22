package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port      string `envconfig:"PORT" default:"8080"`
	DbURL     string `envconfig:"DB_URL" required:"true"`
	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`
}

func LoadConfig() Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}
