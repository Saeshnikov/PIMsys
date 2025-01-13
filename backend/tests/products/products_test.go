package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/products/suite"

	proto "pim-sys/gen/go/products"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func TestProducts_HappyPath(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	respNewProduct, err := st.ProductsClient.NewProduct(ctx, &proto.ProductInfo{
		Name:       gofakeit.FirstName(),
		CategoryId: 1,
		BranchId:   1,
		Status:     "stock",
		Price:      1.0,
		Amount:     30,
		Attributes: []*proto.Attribute{
			{
				Id:        1,
				ValueText: "some text",
			},
		},
	})
	require.NoError(t, err)

	require.NotNil(t, respNewProduct)

	respListProducts, err := st.ProductsClient.ListProducts(ctx, &proto.Empty{})
	require.NoError(t, err)

	respAlterProduct, err := st.ProductsClient.AlterProduct(ctx, &proto.ProductInfoWithId{
		ProductId: 1,
		Product: &proto.ProductInfo{
			Name:       "new-name",
			CategoryId: 1,
			BranchId:   1,
			Status:     "out_of_stock",
			Price:      1.0,
			Amount:     30,
			Attributes: []*proto.Attribute{
				{
					Id:        1,
					ValueText: "new text",
				},
			},
		},
	})
	require.NoError(t, err)

	require.NotNil(t, respAlterProduct)

	respSellProduct, err := st.ProductsClient.SellProduct(ctx, &proto.SellProductRequest{
		ProductId: 1,
		Amount:    10,
	})
	require.NoError(t, err)

	require.NotNil(t, respSellProduct)

	respListProducts, err = st.ProductsClient.ListProducts(ctx, &proto.Empty{})
	require.NoError(t, err)
	require.Equal(t, "new-name", respListProducts.GetProduct()[0].Product.Name)

	respDeleteProducts, err := st.ProductsClient.DeleteProduct(ctx, &proto.DeleteProductRequest{
		ProductId: 1,
	})
	require.NoError(t, err)

	require.NotNil(t, respDeleteProducts)

}
