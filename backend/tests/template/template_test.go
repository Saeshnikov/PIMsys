package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/template/suite"

	proto "pim-sys/gen/go/template"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func DefaultTemplateTest(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	// New template request
	attribute1 := proto.AttributeInfo{
		Description:     "aaa",
		IsUnique:        true,
		IsValueRequired: false,
		Type:            "string",
	}

	var attributes1 []*proto.AttributeInfo
	attributes1 = append(attributes1, &attribute1)

	respNewTemplate, err := st.TemplateClient.NewTemplate(ctx, &proto.NewTemplateRequest{
		Name:        gofakeit.Name(),
		Description: gofakeit.LastName(),
		BranchId:    1,
		Attributes:  attributes1,
	})

	require.NoError(t, err)
	require.NotNil(t, respNewTemplate)

	// Alter template request
	attribute2 := proto.AttributeInfo{
		Description:     "aaa2",
		IsUnique:        true,
		IsValueRequired: false,
		Type:            "string",
	}

	var attributes2 []*proto.AttributeInfo
	attributes2 = append(attributes2, &attribute2)

	respAlterTemplate, err := st.TemplateClient.AlterTemplate(ctx, &proto.AlterTemplateRequest{
		Name:        gofakeit.Name(),
		Description: gofakeit.LastName(),
		BranchId:    1,
		Attributes:  attributes2,
	})

	require.NoError(t, err)
	require.NotNil(t, respAlterTemplate)

	// List templates request
	respListShops, err := st.TemplateClient.ListTemplates(ctx, &proto.ListTemplatesRequest{})
	require.NoError(t, err)
	require.Equal(t, "new-name", respListShops.GetInfo()[0].Name)

	// Delete template request
	respDeleteShop, err := st.TemplateClient.DeleteTemplate(ctx, &proto.DeleteTemplateRequest{
		TemplateId: 1,
	})

	require.NoError(t, err)
	require.NotNil(t, respDeleteShop)

}
