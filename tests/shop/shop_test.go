package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/shop/suite"

	proto "pim-sys/gen/go/shop"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	token := assistance.GetAccessToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	resp, err := st.ShopClient.NewShop(ctx, &proto.NewShopRequest{
		Name:        gofakeit.FirstName(),
		Description: gofakeit.LastName(),
		Url:         gofakeit.URL(),
	})
	require.NoError(t, err)

	require.Equal(t, resp.Message, "alala")

	response, err := st.ShopClient.AlterShop(ctx, &proto.AlterShopRequest{
		ShopId: gofakeit.Int32(),
		ShopInfo: &proto.NewShopRequest{
			Name:        gofakeit.FirstName(),
			Description: gofakeit.LastName(),
			Url:         gofakeit.URL(),
		},
	})
	require.NoError(t, err)

	require.Equal(t, response.Message, "alala")
}
