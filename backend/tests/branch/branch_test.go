package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/branch/suite"

	proto "pim-sys/gen/go/branch"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
	types := [3]string{"online", "marketplace", "offline"}
	respNewShop, err := st.BranchClient.NewBranch(ctx, &proto.NewBranchRequest{
		Name:        gofakeit.FirstName(),
		ShopId:      1,
		Description: gofakeit.LastName(),
		Address:     gofakeit.Address().Address,
		Site:        gofakeit.URL(),
		BranchType:  types[rand.Intn(2)+1],
	})
	require.NoError(t, err)

	require.NotNil(t, respNewShop)

	respAlterShop, err := st.BranchClient.AlterBranch(ctx, &proto.AlterBranchRequest{
		BranchId: 2,
		BranchInfo: &proto.BranchInfo{
			Name:        "new-name",
			Description: "new-description",
			Address:     "new-address",
			Site:        "new-site",
		},
	})
	require.NoError(t, err)

	require.NotNil(t, respAlterShop)

	respListShops, err := st.BranchClient.ListBranches(ctx, &proto.ListBranchesRequest{
		ShopId: 1,
	})
	require.NoError(t, err)
	require.Equal(t, "new-name", respListShops.GetInfo()[1].Name)

	respDeleteShop, err := st.BranchClient.DeleteBranch(ctx, &proto.DeleteBranchRequest{
		BranchId: 2,
	})
	require.NoError(t, err)

	require.NotNil(t, respDeleteShop)

}

func TestAlterBranchWithIncorrectBranchId(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	_, err := st.BranchClient.AlterBranch(ctx, &proto.AlterBranchRequest{
		BranchId: 100,
		BranchInfo: &proto.BranchInfo{
			Name:        "new-name",
			Description: "new-description",
			Address:     "new-address",
			Site:        "new-site",
		},
	})
	require.Error(t, err)
}

func TestListBranchesWithIncorrectShopId(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	list, _ := st.BranchClient.ListBranches(ctx, &proto.ListBranchesRequest{
		ShopId: 100,
	})
	require.Nil(t, list.Info)
}

func TestDeleteBranchWithIncorrectBranchId(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	_, err := st.BranchClient.DeleteBranch(ctx, &proto.DeleteBranchRequest{
		BranchId: 100,
	})
	require.Error(t, err)
}
