package models

import "github.com/golang-jwt/jwt/v4"

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	Uid      int64  `json:"uid" db:"uid"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
