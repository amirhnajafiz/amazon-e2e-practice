package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Auth handles JWT token generation and parsing.
type Auth struct {
	key    string
	expire int
}

// New builds a new auth struct for handling JWT tokens.
func New(pk string, epx int) *Auth {
	return &Auth{
		key:    pk,
		expire: epx,
	}
}

// GenerateJWT creates a new JWT token.
func (a *Auth) GenerateJWT(username string) (string, error) {
	// create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// create claims
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(a.expire)).Unix()
	claims["authorized"] = true

	// generate token string
	tokenString, err := token.SignedString([]byte(a.key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
