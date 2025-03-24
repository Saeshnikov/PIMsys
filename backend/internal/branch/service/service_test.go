package branch_service_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/branch"
	auth_service "pim-sys/internal/branch/service"
	branch_service "pim-sys/internal/branch/service"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//14 tests
//go:generate mockery --testonly --name Branch --dir . --outpkg branch_service_test --output .

func TestServerAPI_NewBranch(t *testing.T) {
	type fields struct {
		UnimplementedBranchServer proto.UnimplementedBranchServer
		Branch                    func() *Branch
	}
	type args struct {
		ctx context.Context
		in  *proto.NewBranchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.NewBranchResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"NewBranch",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).Return(nil)
					return branch
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "test-name",
					ShopId:      1,
					Description: "test-description",
					Address:     "test-address",
					Site:        "test-site",
					BranchType:  "online",
				},
			},
			wantErr: false,
		},
		{
			name: "name is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "",
					ShopId:      1,
					Description: "test-description",
					Address:     "test-address",
					Site:        "test-site",
					BranchType:  "online",
				},
			},
			wantErr: true,
		},
		{
			name: "shopid is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "test-name",
					ShopId:      0,
					Description: "test-description",
					Address:     "test-address",
					Site:        "test-site",
					BranchType:  "online",
				},
			},
			wantErr: true,
		},
		{
			name: "description is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "test-name",
					ShopId:      1,
					Description: "",
					Address:     "test-address",
					Site:        "test-site",
					BranchType:  "online",
				},
			},
			wantErr: true,
		},
		{
			name: "address or site is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "test-name",
					ShopId:      1,
					Description: "test-description",
					Address:     "",
					Site:        "",
					BranchType:  "online",
				},
			},
			wantErr: true,
		},
		{
			name: "some db error",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"NewBranch",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).Return(errors.New("some branch error"))
					return branch
				},
			},
			args: args{
				in: &proto.NewBranchRequest{
					Name:        "test-name",
					ShopId:      1,
					Description: "test-description",
					Address:     "test-address",
					Site:        "test-site",
					BranchType:  "online",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &auth_service.ServerAPI{
				UnimplementedBranchServer: tt.fields.UnimplementedBranchServer,
				Branch:                    tt.fields.Branch(),
			}
			_, err := s.NewBranch(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_AlterBranch(t *testing.T) {
	type fields struct {
		UnimplementedBranchServer proto.UnimplementedBranchServer
		Branch                    func() *Branch
	}
	type args struct {
		ctx context.Context
		in  *proto.AlterBranchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.AlterBranchResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"AlterBranch",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).Return(nil)
					return branch
				},
			},
			args: args{
				in: &proto.AlterBranchRequest{
					BranchId: 1,
					BranchInfo: &proto.BranchInfo{
						BranchId:    1,
						Name:        "test-name",
						ShopId:      1,
						Description: "test-description",
						Address:     "test-address",
						Site:        "test-site",
						BranchType:  "online",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "branchid is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.AlterBranchRequest{
					BranchId: 0,
					BranchInfo: &proto.BranchInfo{
						BranchId:    0,
						Name:        "test-name",
						ShopId:      1,
						Description: "test-description",
						Address:     "test-address",
						Site:        "test-site",
						BranchType:  "online",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "alter branch error",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"AlterBranch",
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
						mock.Anything,
					).Return(errors.New("some db error"))
					return branch
				},
			},
			args: args{
				in: &proto.AlterBranchRequest{
					BranchId: 1,
					BranchInfo: &proto.BranchInfo{
						BranchId:    1,
						Name:        "test-name",
						ShopId:      1,
						Description: "test-description",
						Address:     "test-address",
						Site:        "test-site",
						BranchType:  "online",
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &auth_service.ServerAPI{
				UnimplementedBranchServer: tt.fields.UnimplementedBranchServer,
				Branch:                    tt.fields.Branch(),
			}
			_, err := s.AlterBranch(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_DeleteBranch(t *testing.T) {
	type fields struct {
		UnimplementedBranchServer proto.UnimplementedBranchServer
		Branch                    func() *Branch
	}
	type args struct {
		ctx context.Context
		in  *proto.DeleteBranchRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.DeleteBranchResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"DeleteBranch",
						mock.Anything,
						mock.Anything,
					).Return(nil)
					return branch
				},
			},
			args: args{
				in: &proto.DeleteBranchRequest{
					BranchId: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "branchid is required",
			fields: fields{
				Branch: func() *Branch {
					return &Branch{}
				},
			},
			args: args{
				in: &proto.DeleteBranchRequest{
					BranchId: 0,
				},
			},
			wantErr: true,
		},
		{
			name: "delete branch error",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"DeleteBranch",
						mock.Anything,
						mock.Anything,
					).Return(errors.New("some db error"))
					return branch
				},
			},
			args: args{
				in: &proto.DeleteBranchRequest{
					BranchId: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branch_service.ServerAPI{
				UnimplementedBranchServer: tt.fields.UnimplementedBranchServer,
				Branch:                    tt.fields.Branch(),
			}
			_, err := s.DeleteBranch(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_ListBranches(t *testing.T) {
	type fields struct {
		UnimplementedBranchServer proto.UnimplementedBranchServer
		Branch                    func() *Branch
	}
	type args struct {
		ctx context.Context
		in  *proto.ListBranchesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *proto.ListBranchesResponse
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"ListBranches",
						mock.Anything,
						mock.Anything,
					).Return(nil, nil)
					return branch
				},
			},
			args: args{
				in: &proto.ListBranchesRequest{
					ShopId: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "list branches error",
			fields: fields{
				Branch: func() *Branch {
					branch := &Branch{}
					branch.Mock.On(
						"ListBranches",
						mock.Anything,
						mock.Anything,
					).Return(nil, errors.New("some db error"))
					return branch
				},
			},
			args: args{
				in: &proto.ListBranchesRequest{
					ShopId: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &branch_service.ServerAPI{
				UnimplementedBranchServer: tt.fields.UnimplementedBranchServer,
				Branch:                    tt.fields.Branch(),
			}
			_, err := s.ListBranches(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
