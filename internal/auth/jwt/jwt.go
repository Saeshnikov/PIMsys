package auth_jwt

import (
	"time"

	"pim-sys/internal/auth/storage"

	"github.com/golang-jwt/jwt/v5"
)

// NewToken creates new JWT token for given user and app.
func NewToken(user storage.User, app storage.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
