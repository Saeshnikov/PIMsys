package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/template/suite"

	proto "pim-sys/gen/go/template"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func DefaultTemplateTest(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	respNewTemplate, err := st.TemplateClient.NewTemplate(ctx, &proto.NewTemplateRequest{
		// Name:        gofakeit.FirstName(),
		// Description: gofakeit.LastName(),
		// Url:         gofakeit.URL(),
	})

	require.NoError(t, err)
	require.NotNil(t, respNewTemplate)

	respAlterTemplate, err := st.TemplateClient.AlterTemplate(ctx, &proto.AlterTemplateRequest{
		// ShopId: 1,
		// ShopInfo: &proto.ShopInfo{
		// 	Name:        "new-name",
		// 	Description: "new-description",
		// 	Url:         "new-url",
		// },
	})

	require.NoError(t, err)
	require.NotNil(t, respAlterTemplate)

	respListShops, err := st.TemplateClient.ListTemplates(ctx, &proto.ListTemplatesRequest{})
	require.NoError(t, err)
	require.Equal(t, "new-name", respListShops.GetInfo()[0].Name)

	respDeleteShop, err := st.TemplateClient.DeleteTemplate(ctx, &proto.DeleteTemplateRequest{
		// ShopId: 1,
	})

	require.NoError(t, err)
	require.NotNil(t, respDeleteShop)

}
