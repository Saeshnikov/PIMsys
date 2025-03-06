package template_service_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/template"
	template_service "pim-sys/internal/template/service"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Template --dir . --outpkg template_service_test --output .

// 13 tests
func TestServerAPI_NewTemplate(t *testing.T) {
	type fields struct {
		UnimplementedTemplateServer proto.UnimplementedTemplateServer
		template                    func() template_service.Template
	}
	type args struct {
		ctx context.Context
		in  *proto.NewTemplateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.NewTemplateResponse
		wantErr bool
	}{
		{
			name: "NewTemplate good",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "description",
					BranchId:    1,
					Attributes: []*proto.AttributeInfo{
						{}, {}, {}, {}, {},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "NewTemplate bad name",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "",
					Description: "description",
					BranchId:    1,
					Attributes: []*proto.AttributeInfo{
						{}, {}, {}, {}, {},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "NewTemplate bad description",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "",
					BranchId:    1,
					Attributes: []*proto.AttributeInfo{
						{}, {}, {}, {}, {},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "NewTemplate bad branchId",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "description",
					BranchId:    0,
					Attributes: []*proto.AttributeInfo{
						{}, {}, {}, {}, {},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "NewTemplate bad attribute",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "description",
					BranchId:    1,
					Attributes:  []*proto.AttributeInfo{},
				},
			},
			wantErr: true,
		},
		{
			name: "NewTemplate attribute more than 10",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "description",
					BranchId:    1,
					Attributes:  []*proto.AttributeInfo{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}, {}},
				},
			},
			wantErr: true,
		},
		{
			name: "NewTemplate error on NewTemplate",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"NewTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(errors.New("error!!!"))
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.NewTemplateRequest{
					Name:        "name",
					Description: "description",
					BranchId:    1,
					Attributes:  []*proto.AttributeInfo{{}},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &template_service.ServerAPI{
				UnimplementedTemplateServer: tt.fields.UnimplementedTemplateServer,
				Template:                    tt.fields.template(),
			}
			_, err := s.NewTemplate(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerAPI.NewTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestServerAPI_DeleteTemplate(t *testing.T) {
	type fields struct {
		UnimplementedTemplateServer proto.UnimplementedTemplateServer
		template                    func() template_service.Template
	}
	type args struct {
		ctx context.Context
		in  *proto.DeleteTemplateRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.DeleteTemplateResponse
		wantErr bool
	}{
		{
			name: "DeleteTemplate default",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.DeleteTemplateRequest{
					TemplateId: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "DeleteTemplate invalid templateId",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything).Return(nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.DeleteTemplateRequest{
					TemplateId: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "DeleteTemplate error on DeleteTemplate",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything).Return(errors.New("Error!"))
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.DeleteTemplateRequest{
					TemplateId: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &template_service.ServerAPI{
				UnimplementedTemplateServer: tt.fields.UnimplementedTemplateServer,
				Template:                    tt.fields.template(),
			}
			_, err := s.DeleteTemplate(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerAPI.DeleteTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestServerAPI_ListTemplates(t *testing.T) {
	type fields struct {
		UnimplementedTemplateServer proto.UnimplementedTemplateServer
		template                    func() template_service.Template
	}
	type args struct {
		ctx context.Context
		in  *proto.ListTemplatesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ListTemplatesResponse
		wantErr bool
	}{
		{
			name: "ListTemplates default",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"ListTemplates",
						mock.Anything,
						mock.Anything).Return([]*proto.TemplateInfo{}, nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.ListTemplatesRequest{
					BranchId: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "ListTemplates invalid branchId",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"ListTemplates",
						mock.Anything,
						mock.Anything).Return([]*proto.TemplateInfo{}, nil)
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.ListTemplatesRequest{
					BranchId: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "ListTemplates invalid branchId",
			fields: fields{
				template: func() template_service.Template {
					template := &Template{}
					template.Mock.On(
						"ListTemplates",
						mock.Anything,
						mock.Anything).Return([]*proto.TemplateInfo{}, errors.New("Error!"))
					return template
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				in: &proto.ListTemplatesRequest{
					BranchId: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &template_service.ServerAPI{
				UnimplementedTemplateServer: tt.fields.UnimplementedTemplateServer,
				Template:                    tt.fields.template(),
			}
			_, err := s.ListTemplates(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ServerAPI.ListTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
