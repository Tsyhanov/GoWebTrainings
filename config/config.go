package config

import (
	"log"
	"os"
	"test-registration-form/models"

	"github.com/labstack/echo/v4/middleware"
)

type AppConfig struct {
	Port      string
	JWTSecret string
}

var Config AppConfig
var AuthConfig middleware.JWTConfig

func SetConfig() {
	Config = AppConfig{
		/*
			Port:      getEnv("port", "5000"),
			JWTSecret: getEnv("jwtSecret", ""),
		*/
		Port:      ":8080",
		JWTSecret: "my_private_key",
	}

	AuthConfig = middleware.JWTConfig{
		Claims:      &models.JWTClaims{},
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:Authorization",
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else if defaultVal == "" {
		log.Fatalf("environment variable %s cannot have a nil value", key)
	}
	return defaultVal
}
