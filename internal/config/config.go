package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name string
		Port string
		Mode string
	}
	Redis struct {
		Host string
		Port string
	}
	JWT struct {
		SecretKey string
	}
}

var instance *Config = nil

func GetInstance() *Config {
	return instance
}

func Load() (*Config, error) {
	if instance != nil {
		return instance, nil
	}

	instance = &Config{}

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	instance.App.Name = os.Getenv("APP_NAME")
	instance.App.Port = os.Getenv("APP_PORT")
	instance.App.Mode = os.Getenv("APP_MODE")

	instance.Redis.Host = os.Getenv("REDIS_HOST")
	instance.Redis.Port = os.Getenv("REDIS_PORT")

	instance.JWT.SecretKey = os.Getenv("JWT_SECRET")

	return instance, nil
}
