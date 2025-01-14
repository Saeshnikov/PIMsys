package tests

import (
	"testing"
	"time"

	proto "pim-sys/gen/go/sso"
	suite "pim-sys/tests/sso/suite"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	appSecret = "test-secret"

	passDefaultLen = 10
	configPath     = "suite/config.yaml"
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t, configPath)

	email := gofakeit.FirstName()
	pass := randomFakePassword()
	name := gofakeit.FirstName()
	phone := gofakeit.Phone()

	respReg, err := st.AuthClient.Register(ctx, &proto.RegisterRequest{
		Email:    email,
		Password: pass,
		Name:     name,
		Phone:    phone,
	})
	require.NoError(t, err)
	assert.NotEmpty(t, respReg.GetUserId())

	respLogin, err := st.AuthClient.Login(ctx, &proto.LoginRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)

	token := respLogin.GetToken()
	require.NotEmpty(t, token)

	loginTime := time.Now()

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(appSecret), nil
	})
	require.NoError(t, err, tokenParsed)

	claims, ok := tokenParsed.Claims.(jwt.MapClaims)
	require.True(t, ok)

	assert.Equal(t, respReg.GetUserId(), int64(claims["uid"].(float64)))
	assert.Equal(t, email, claims["email"].(string))

	const deltaSeconds = 1

	// check if exp of token is in correct range, ttl get from st.Cfg.TokenTTL
	assert.InDelta(t, loginTime.Add(st.Cfg.TokenTLL).Unix(), claims["exp"].(float64), deltaSeconds)
}

func TestRegisterLogin_DuplicatedRegistration(t *testing.T) {
	ctx, st := suite.New(t, configPath)

	email := gofakeit.FirstName()
	pass := randomFakePassword()
	name := gofakeit.FirstName()
	phone := gofakeit.Phone()

	respReg, err := st.AuthClient.Register(ctx, &proto.RegisterRequest{
		Email:    email,
		Password: pass,
		Name:     name,
		Phone:    phone,
	})
	require.NoError(t, err)
	require.NotEmpty(t, respReg.GetUserId())

	respReg, err = st.AuthClient.Register(ctx, &proto.RegisterRequest{
		Email:    email,
		Password: pass,
		Name:     name,
		Phone:    phone,
	})
	require.Error(t, err)
	assert.Empty(t, respReg.GetUserId())

	assert.ErrorContains(t, err, "user already exists")
}

func TestRegister_FailCases(t *testing.T) {
	ctx, st := suite.New(t, configPath)

	tests := []struct {
		name        string
		email       string
		password    string
		nameuser    string
		phone       string
		expectedErr string
	}{
		{
			name:        "Register with Empty Password",
			email:       gofakeit.FirstName(),
			password:    "",
			nameuser:    gofakeit.FirstName(),
			phone:       gofakeit.Phone(),
			expectedErr: "password is required",
		},
		{
			name:        "Register with Empty Email",
			email:       "",
			password:    randomFakePassword(),
			nameuser:    gofakeit.FirstName(),
			phone:       gofakeit.Phone(),
			expectedErr: "email is required",
		},
		{
			name:        "Register with Both Empty",
			email:       "",
			password:    "",
			nameuser:    gofakeit.FirstName(),
			phone:       gofakeit.Phone(),
			expectedErr: "email is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &proto.RegisterRequest{
				Email:    tt.email,
				Password: tt.password,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)

		})
	}
}

func TestLogin_FailCases(t *testing.T) {
	ctx, st := suite.New(t, configPath)

	tests := []struct {
		name        string
		email       string
		password    string
		expectedErr string
	}{
		{
			name:        "Login with Empty Password",
			email:       gofakeit.FirstName(),
			password:    "",
			expectedErr: "password is required",
		},
		{
			name:        "Login with Empty Email",
			email:       "",
			password:    randomFakePassword(),
			expectedErr: "email is required",
		},
		{
			name:        "Login with Both Empty Email and Password",
			email:       "",
			password:    "",
			expectedErr: "email is required",
		},
		{
			name:        "Login with Non-Matching Password",
			email:       gofakeit.FirstName(),
			password:    randomFakePassword(),
			expectedErr: "invalid email or password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.AuthClient.Register(ctx, &proto.RegisterRequest{
				Email:    gofakeit.FirstName(),
				Password: randomFakePassword(),
				Name:     gofakeit.FirstName(),
				Phone:    gofakeit.Phone(),
			})
			require.NoError(t, err)

			_, err = st.AuthClient.Login(ctx, &proto.LoginRequest{
				Email:    tt.email,
				Password: tt.password,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, false, false, false, passDefaultLen)
}
