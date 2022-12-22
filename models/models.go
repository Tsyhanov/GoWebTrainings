package models

import (
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	ID       int64  `gorm:"primary_key" json:"id"`
	Email    string `gorm:"not null;unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Name     string `gorm:"type:varchar(100);not null;unique" json:"name" form:"name"`
}

type Post struct {
	Id   int64  `json:"uid" gorm:"id"`
	Body string `json:"body" gorm:"body"`
}

type Comment struct {
	Id     int64  `json:"uid" gorm:"id"`
	Body   string `json:"body" gorm:"body"`
	PostId int64  `json:"postid" gorm:"postid"`
}
