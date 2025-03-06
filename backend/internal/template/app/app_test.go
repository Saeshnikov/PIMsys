package template_app_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/template"
	template_app "pim-sys/internal/template/app"
	"reflect"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Storage --dir . --outpkg template_app_test --output .

// 9 tests

func TestTemplate_NewTemplate(t *testing.T) {
	type fields struct {
		templateStorage func() template_app.Storage
	}
	type args struct {
		ctx         context.Context
		name        string
		description string
		branch_id   int32
		attributes  []*proto.AttributeInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"CreateTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					return storage
				},
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				name:        "asd",
				description: "asdad",
				branch_id:   1,
				attributes:  nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := &template_app.Template{
				TemplateStorage: tt.fields.templateStorage(),
			}
			if err := template.NewTemplate(tt.args.ctx, tt.args.name, tt.args.description, tt.args.branch_id, tt.args.attributes); (err != nil) != tt.wantErr {
				t.Errorf("Template.NewTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemplate_DeleteTemplate(t *testing.T) {
	type fields struct {
		templateStorage func() template_app.Storage
	}
	type args struct {
		ctx        context.Context
		templateId int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "userHaveAccessToTemplate",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(5), nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2, 5}, nil)
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: false,
		},
		{
			name: "userDontHaveAccessToTemplate",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(5), nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2, 3}, nil)
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: true,
		},
		{
			name: "DeleteTemplate error return test",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(1), errors.New("error!!!"))
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2, 3}, nil)
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: true,
		},
		{
			name: "userMustHaveAccess error user_id cast to int32",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(1), nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2, 3}, nil)
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "invalid"})),
			},
			wantErr: true,
		},
		{
			name: "userMustHaveAccess error context",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(1), nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2, 3}, nil)
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id1": "1"})),
			},
			wantErr: true,
		},
		{
			name: "userMustHaveAccess error userListBranches",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"DeleteTemplate",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything).Return(nil)
					storage.Mock.On(
						"GetBranchIdFromTemplateId",
						mock.Anything,
						mock.Anything).Return(int32(1), nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return(nil, errors.New("error!!!"))
					return storage
				},
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := &template_app.Template{
				TemplateStorage: tt.fields.templateStorage(),
			}
			if err := template.DeleteTemplate(tt.args.ctx, tt.args.templateId); (err != nil) != tt.wantErr {
				t.Errorf("Template.DeleteTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTemplate_ListTemplates(t *testing.T) {
	type fields struct {
		templateStorage func() template_app.Storage
	}
	type args struct {
		ctx       context.Context
		branch_id int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*proto.TemplateInfo
		wantErr bool
	}{
		{
			name: "ListTemplateUserHaveAccessToTemplate",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"ListTemplates",
						mock.Anything,
						mock.Anything).Return(nil, nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{1, 2}, nil)
					return storage
				},
			},
			args: args{
				ctx:       metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: false,
		},
		{
			name: "ListTemplateUserDontHaveAccessToTemplate",
			fields: fields{
				templateStorage: func() template_app.Storage {
					storage := &Storage{}
					storage.Mock.On(
						"ListTemplates",
						mock.Anything,
						mock.Anything).Return(nil, nil)
					storage.Mock.On(
						"GetUserListBranches",
						mock.Anything,
						mock.Anything).Return([]int32{3, 2}, nil)
					return storage
				},
			},
			args: args{
				ctx:       metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branch_id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			template := &template_app.Template{
				TemplateStorage: tt.fields.templateStorage(),
			}
			got, err := template.ListTemplates(tt.args.ctx, tt.args.branch_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.ListTemplates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.ListTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
