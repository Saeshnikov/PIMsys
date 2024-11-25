package assistance

import (
	"testing"
	"time"

	auth_jwt "pim-sys/internal/auth/jwt"
	"pim-sys/internal/auth/storage"

	"github.com/stretchr/testify/require"
)

func GetTestToken(t *testing.T) string {
	usr := storage.User{
		ID:       1,
		Name:     "test-name",
		Email:    "test-email",
		PassHash: []byte("test-hash"),
		Phone:    "test-phone",
		IsAdmin:  false,
	}

	token, err := auth_jwt.NewToken(usr, time.Hour)
	require.NoError(t, err)

	return token
}
