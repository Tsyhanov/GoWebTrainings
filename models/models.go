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
	ID     int64  `json:"id" gorm:"id"`
	UserId int64  `json:"userId" gorm:"userid"`
	Title  string `json:"title" gorm:"title"`
	Body   string `json:"body" gorm:"body"`
}

type Comment struct {
	ID     int64  `json:"id" gorm:"id"`
	Email  string `json:"email" gorm:"email"`
	Name   string `json:"name" gorm:"name"`
	Body   string `json:"body" gorm:"body"`
	PostId int64  `json:"postId" gorm:"postid"` //foreign key
	Post   Post
}
