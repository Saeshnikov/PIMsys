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
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	respNewShop, err := st.ShopClient.NewShop(ctx, &proto.NewShopRequest{
		Name:        gofakeit.FirstName(),
		Description: gofakeit.LastName(),
		Url:         gofakeit.URL(),
	})
	require.NoError(t, err)

	require.NotNil(t, respNewShop)

	respAlterShop, err := st.ShopClient.AlterShop(ctx, &proto.AlterShopRequest{
		ShopId: 2,
		ShopInfo: &proto.ShopInfo{
			Name:        "new-name",
			Description: "new-description",
			Url:         "new-url",
		},
	})
	require.NoError(t, err)

	require.NotNil(t, respAlterShop)

	respListShops, err := st.ShopClient.ListShops(ctx, &proto.ListShopsRequest{})
	require.NoError(t, err)
	require.Equal(t, "new-name", respListShops.GetInfo()[1].Name)

	respDeleteShop, err := st.ShopClient.DeleteShop(ctx, &proto.DeleteShopRequest{
		ShopId: 2,
	})
	require.NoError(t, err)

	require.NotNil(t, respDeleteShop)

}
