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

func TestTemplateNew(t *testing.T) {
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
}

func TestTemplateList(t *testing.T) {
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
	listResp, err := st.TemplateClient.ListTemplates(ctx, &proto.ListTemplatesRequest{
		BranchId: 1,
	})
	require.NoError(t, err)

	require.Equal(t, name, listResp.Info[len(listResp.Info)-1].Name)
}

func TestTemplateDelete(t *testing.T) {
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

	// Delete template request
	respDeleteTemplate, err := st.TemplateClient.DeleteTemplate(ctx, &proto.DeleteTemplateRequest{
		TemplateId: 1,
	})

	require.NoError(t, err)
	require.NotNil(t, respDeleteTemplate)
}

func TestTemplateNewIncorrectBranchID(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	// Create with full attribute info
	attr := &proto.AttributeInfo{
		Name:            "user_email",
		Type:            "text",
		IsUnique:        true,
		IsValueRequired: true,
		Description:     "Primary user contact",
	}

	uniqueName := gofakeit.Name()
	_, err := st.TemplateClient.NewTemplate(ctx, &proto.NewTemplateRequest{
		Name:       uniqueName,
		BranchId:   10,
		Attributes: []*proto.AttributeInfo{attr},
	})
	require.Error(t, err)
}

func TestTemplateDeleteIncorrectTemplateID(t *testing.T) {
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

	// incorrect Delete template request
	_, err = st.TemplateClient.DeleteTemplate(ctx, &proto.DeleteTemplateRequest{
		TemplateId: 100,
	})

	require.Error(t, err)
}

func TestTemplateListIncorrectBranchID(t *testing.T) {
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
		BranchId: 100,
	})
	require.Error(t, err)
}
