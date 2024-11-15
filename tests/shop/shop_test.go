package tests

import (
	"testing"

	"pim-sys/tests/shop/suite"

	proto "pim-sys/gen/go/shop"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)

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
