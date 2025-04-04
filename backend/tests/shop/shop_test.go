package tests

import (
	"context"
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/shop/suite"

	proto "pim-sys/gen/go/shop"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func TestShop_HappyPath(t *testing.T) {
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

func TestShop_AlterShop_IncorrectShopId(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	respAlterShop, err := st.ShopClient.AlterShop(ctx, &proto.AlterShopRequest{
		ShopId: 99,
		ShopInfo: &proto.ShopInfo{
			Name:        "new-name",
			Description: "new-description",
			Url:         "new-url",
		},
	})
	require.Nil(t, respAlterShop)
	require.Error(t, err)
}

func TestShop_DeleteShop_IncorrectShopId(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	respDeleteShop, err := st.ShopClient.DeleteShop(ctx, &proto.DeleteShopRequest{
		ShopId: 99,
	})

	require.Nil(t, respDeleteShop)
	require.Error(t, err)
}

func TestShop_Unauthentication(t *testing.T) {
	ctx, st := suite.New(t, configPath)
	ctx = context.TODO()

	respNewShop, err := st.ShopClient.NewShop(ctx, &proto.NewShopRequest{
		Name:        gofakeit.FirstName(),
		Description: gofakeit.LastName(),
		Url:         gofakeit.URL(),
	})
	require.Error(t, err)

	require.Nil(t, respNewShop)

	respAlterShop, err := st.ShopClient.AlterShop(ctx, &proto.AlterShopRequest{
		ShopId: 2,
		ShopInfo: &proto.ShopInfo{
			Name:        "new-name",
			Description: "new-description",
			Url:         "new-url",
		},
	})
	require.Error(t, err)

	require.Nil(t, respAlterShop)

	respListShops, err := st.ShopClient.ListShops(ctx, &proto.ListShopsRequest{})
	require.Error(t, err)
	require.Nil(t, respListShops)

	respDeleteShop, err := st.ShopClient.DeleteShop(ctx, &proto.DeleteShopRequest{
		ShopId: 2,
	})
	require.Error(t, err)
	require.Nil(t, respDeleteShop)

}
