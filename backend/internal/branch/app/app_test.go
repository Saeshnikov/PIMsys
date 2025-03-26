package branch_app_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/branch"
	branch_app "pim-sys/internal/branch/app"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

//11 tests
//go:generate mockery --testonly --name Storage --dir . --outpkg branch_app_test --output .

func TestBranch_NewBranch(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		shop_id     int32
		description string
		address     string
		site        string
		branch_type string
	}
	tests := []struct {
		name    string
		branch  func() *branch_app.Branch
		args    args
		wantErr bool
	}{
		{
			name: "deafult",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"CreateBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				name:        "test-name",
				shop_id:     1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: false,
		},
		{
			name: "invalid meta",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"CreateBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "test-name",
				shop_id:     1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: true,
		},
		{
			name: "invalid user_id",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"CreateBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "invalid"})),
				name:        "test-name",
				shop_id:     1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: true,
		},
		{
			name: "access denied",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"CreateBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				name:        "test-name",
				shop_id:     1,
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
				branch_type: "online",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.branch().NewBranch(tt.args.ctx, tt.args.name, tt.args.shop_id, tt.args.description, tt.args.address, tt.args.site, tt.args.branch_type)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

		})
	}
}

func TestBranch_AlterBranch(t *testing.T) {
	type fields struct {
		BranchStorage Storage
	}
	type args struct {
		ctx         context.Context
		branchId    int32
		name        string
		description string
		address     string
		site        string
	}
	tests := []struct {
		name    string
		branch  func() *branch_app.Branch
		args    args
		wantErr bool
	}{
		{
			name: "deafult",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(1), nil)
				storage.Mock.On(
					"AlterBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId:    1,
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
			},
			wantErr: false,
		},
		{
			name: "failed to get shopid",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(0), errors.New("some error"))
				storage.Mock.On(
					"AlterBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId:    1,
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
			},
			wantErr: true,
		},
		{
			name: "failed permission",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(1), nil)
				storage.Mock.On(
					"AlterBranch",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId:    1,
				name:        "test-name",
				description: "test-description",
				address:     "test-address",
				site:        "test-site",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.branch().AlterBranch(tt.args.ctx, tt.args.branchId, tt.args.name, tt.args.description, tt.args.address, tt.args.site)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestBranch_DeleteBranch(t *testing.T) {
	type fields struct {
		BranchStorage Storage
	}
	type args struct {
		ctx      context.Context
		branchId int32
	}
	tests := []struct {
		name    string
		branch  func() *branch_app.Branch
		args    args
		wantErr bool
	}{
		{
			name: "deafult",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(1), nil)
				storage.Mock.On(
					"DeleteBranch",
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:      metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId: 1,
			},
			wantErr: false,
		},
		{
			name: "failed to get shopid",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(0), errors.New("some error"))
				storage.Mock.On(
					"DeleteBranch",
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:      metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId: 1,
			},
			wantErr: true,
		},
		{
			name: "failed permission",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"GetShopId",
					mock.Anything,
					mock.Anything).Return(int32(1), nil)
				storage.Mock.On(
					"DeleteBranch",
					mock.Anything,
					mock.Anything).Return(nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:      metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				branchId: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.branch().DeleteBranch(tt.args.ctx, tt.args.branchId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)

		})
	}
}

func TestBranch_ListBranches(t *testing.T) {
	type fields struct {
		BranchStorage Storage
	}
	type args struct {
		ctx     context.Context
		shop_id int32
	}
	tests := []struct {
		name    string
		branch  func() *branch_app.Branch
		args    args
		want    []*proto.BranchInfo
		wantErr bool
	}{
		{
			name: "deafult",
			branch: func() *branch_app.Branch {
				storage := &Storage{}
				storage.Mock.On(
					"ListBranches",
					mock.Anything,
					mock.Anything).Return([]*proto.BranchInfo{}, nil)
				return &branch_app.Branch{
					BranchStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				shop_id: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.branch().ListBranches(tt.args.ctx, tt.args.shop_id)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
