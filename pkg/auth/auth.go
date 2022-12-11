package auth

import (
	"fmt"
	"net/http"
	"test-registration-form/config"
	"test-registration-form/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GenerateTokensAndSetCookies(user *models.User, c echo.Context) error {
	accessToken, exp, err := generateToken(user, []byte(config.GetJWTSecret()))
	if err != nil {
		return err
	}

	setTokenCookie(config.Config.TokenCookieName, accessToken, exp, c)

	return nil
}

func generateToken(user *models.User, secret []byte) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.JWTClaims{
		Email: "test@test.com",
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	token, err := rawToken.SignedString(secret)
	if err != nil {
		fmt.Println("error:", err)
		return "", time.Now(), err
	}

	return token, expirationTime, nil
}

func setTokenCookie(name, token string, expiration time.Time, c echo.Context) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = token
	cookie.Expires = expiration
	cookie.Path = "/"
	cookie.HttpOnly = true

	c.SetCookie(cookie)
}
