package shop_app_test

import (
	"context"
	proto "pim-sys/gen/go/shop"
	app "pim-sys/internal/shop/app"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Storage --dir . --outpkg shop_app_test --output .

// 12 tests

func TestShop_NewShop(t *testing.T) {
	type args struct {
		ctx         context.Context
		name        string
		description string
		url         string
	}
	tests := []struct {
		name    string
		shop    func() *app.Shop
		args    args
		wantErr bool
	}{
		{
			name: "deafult",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"CreateShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				name:        "asd",
				description: "asdad",
				url:         "asdasd",
			},
			wantErr: false,
		},
		{
			name: "invalid meta",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"CreateShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx:         context.TODO(),
				name:        "asd",
				description: "asdad",
				url:         "asdasd",
			},
			wantErr: true,
		},
		{
			name: "invalid user_id",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"CreateShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx:         metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "invalid"})),
				name:        "asd",
				description: "asdad",
				url:         "asdasd",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.shop().NewShop(tt.args.ctx, tt.args.name, tt.args.description, tt.args.url)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestShop_ListShops(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		shop    func() *app.Shop
		args    args
		want    []*proto.ShopInfo
		wantErr bool
	}{
		{
			name: "deafult",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "invalid meta",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: context.TODO(),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid user_id",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "invalid"})),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.shop().ListShops(tt.args.ctx)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestShop_AlterShop(t *testing.T) {
	type args struct {
		ctx         context.Context
		shopId      int32
		name        string
		description string
		url         string
	}
	tests := []struct {
		name    string
		shop    func() *app.Shop
		args    args
		wantErr bool
	}{
		{
			name: "default",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]*proto.ShopInfo{
					{ShopId: 0},
				}, nil)
				storage.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: false,
		},
		{
			name: "invalid meta",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]*proto.ShopInfo{
					{ShopId: 0},
				}, nil)
				storage.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: context.TODO(),
			},
			wantErr: true,
		},
		{
			name: "invalid user_id",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]*proto.ShopInfo{
					{ShopId: 0},
				}, nil)
				storage.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "invalid"})),
			},
			wantErr: true,
		},
		{
			name: "access denied",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.shop().AlterShop(tt.args.ctx, tt.args.shopId, tt.args.name, tt.args.description, tt.args.url)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestShop_DeleteShop(t *testing.T) {
	type args struct {
		ctx    context.Context
		shopId int32
	}
	tests := []struct {
		name    string
		shop    func() *app.Shop
		args    args
		wantErr bool
	}{
		{
			name: "default",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return([]*proto.ShopInfo{
					{ShopId: 0},
				}, nil)
				storage.Mock.On(
					"DeleteShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: false,
		},
		{
			name: "access_denied",
			shop: func() *app.Shop {
				storage := &Storage{}
				storage.Mock.On(
					"ListShops",
					mock.Anything,
					mock.Anything).Return(nil, nil)
				storage.Mock.On(
					"DeleteShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)
				return &app.Shop{
					ShopStorage: storage,
				}
			},
			args: args{
				ctx: metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.shop().DeleteShop(tt.args.ctx, tt.args.shopId)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
