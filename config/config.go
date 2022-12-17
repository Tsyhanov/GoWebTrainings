package config

import (
	"log"
	"os"
	"test-registration-form/models"

	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/bcrypt"
)

type AppConfig struct {
	Port            string
	TokenCookieName string
	JWTSecret       string
	DBName          string
	DBUser          string
	DBPassword      string
	DBIp            string
	DBPort          string
}

var Config AppConfig
var AuthConfig middleware.JWTConfig

func SetConfig() {
	Config = AppConfig{
		/*
			Port:      getEnv("port", "5000"),
			JWTSecret: getEnv("jwtSecret", ""),
		*/
		Port:            ":8080",
		TokenCookieName: "access-token",
		JWTSecret:       "my_private_key",
		DBName:          "nixdb",
		DBUser:          "root",
		DBPassword:      "weak_password",
		DBIp:            "127.0.0.1",
		DBPort:          "3306",
	}

	AuthConfig = middleware.JWTConfig{
		Claims:      &models.JWTClaims{},
		SigningKey:  []byte(GetJWTSecret()),
		TokenLookup: "cookie:access-token",
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

func GetJWTSecret() string {
	return Config.JWTSecret
}

func LoadTestUser() *models.User {
	// Just for demonstration purpose, we create a user with the encrypted "111" password.
	// In real-world applications, you might load the user from the database by specific parameters (email, username, etc.)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("111"), 8)
	return &models.User{Password: string(hashedPassword), Email: "test@test.com"}
}
