package auth_jwt

import (
	"fmt"
	"time"

	"pim-sys/internal/auth/storage"

	"pim-sys/internal/auth/secret"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user storage.User, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Add(duration).Unix())
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["role"] = user.IsAdmin

	tokenString, err := token.SignedString([]byte(secret.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
