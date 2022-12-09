package models

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
