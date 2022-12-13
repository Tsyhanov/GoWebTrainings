package models

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
