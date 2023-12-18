package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port string `env:"RUN_PORT"`
}

type DatabaseConfig struct {
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbSslMode  string `env:"DB_SSL_MODE"`
}

func ReadConfig() (string, error) {
	var conf Config
	if err := cleanenv.ReadConfig("../../.env", &conf); err != nil {
		return "", err
	}
	return conf.Port, nil
}
