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

func TestTemplateTestDefault(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	// New template request
	attribute1 := proto.AttributeInfo{
		Description:     "aaa",
		IsUnique:        true,
		IsValueRequired: false,
		Type:            "text",
	}

	var attributes1 []*proto.AttributeInfo
	attributes1 = append(attributes1, &attribute1)

	name := gofakeit.Name()
	descr := gofakeit.LastName()
	respNewTemplate, err := st.TemplateClient.NewTemplate(ctx, &proto.NewTemplateRequest{
		Name:        name,
		Description: descr,
		BranchId:    1,
		Attributes:  attributes1,
	})

	require.NoError(t, err)
	require.NotNil(t, respNewTemplate)

	// List templates request
	_, err = st.TemplateClient.ListTemplates(ctx, &proto.ListTemplatesRequest{
		BranchId: 1,
	})
	require.NoError(t, err)

	// var containsTemplInfo bool
	// for _, templInfo := range respListTemplates.GetInfo() {
	// 	if templInfo == proto.TemplateInfo{
	// 		TemplateId:  1,
	// 		Name:        name,
	// 		Description: descr,
	// 		Attributes:  attributes1
	// 	} {

	// 	}
	// }
	// require.Contains(t, respListTemplates.GetInfo(), &proto.TemplateInfo{
	// 	TemplateId:  1,
	// 	Name:        name,
	// 	Description: descr,
	// 	Attributes:  attributes1,
	// })

	// Delete template request
	respDeleteShop, err := st.TemplateClient.DeleteTemplate(ctx, &proto.DeleteTemplateRequest{
		TemplateId: 2,
	})

	require.NoError(t, err)
	require.NotNil(t, respDeleteShop)

}
