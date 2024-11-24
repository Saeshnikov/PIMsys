package assistance

import (
	"fmt"
	proto "pim-sys/gen/go/sso"
	suite "pim-sys/tests/sso/suite"

	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

const configPath = "../sso/suite/config.yaml"

func GetAccessToken(t *testing.T) string {
	ctx, st := suite.New(t, configPath)
	fmt.Println(st.Cfg.Grpc.Port)
	email := gofakeit.Email()
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

	respLog, err := st.AuthClient.Login(ctx, &proto.LoginRequest{
		Email:    email,
		Password: pass,
	})
	require.NoError(t, err)

	token := respLog.GetToken()
	require.NotEmpty(t, token)

	return token
}

func randomFakePassword() string {
	return gofakeit.Password(true, true, true, true, false, 10)
}
