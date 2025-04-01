package shop_service_test

import (
	"context"
	"errors"
	proto "pim-sys/gen/go/shop"
	service "pim-sys/internal/shop/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --testonly --name Shop --dir . --outpkg shop_service_test --output .

// 11 tests

func TestServerAPI_NewShop(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.NewShopRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.NewShopResponse
		wantErr bool
	}{
		{
			name: "correct_1",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"NewShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.NewShopRequest{
					Name:        "name",
					Description: "description",
					Url:         "url",
				},
			},
			want:    &proto.NewShopResponse{},
			wantErr: false,
		},
		{
			name: "correct_2",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"NewShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.NewShopRequest{
					Name:        "name1",
					Description: "description1",
					Url:         "url1",
				},
			},
			want:    &proto.NewShopResponse{},
			wantErr: false,
		},
		{
			name: "correct_3",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"NewShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.NewShopRequest{
					Name:        "1name1",
					Description: "1description1",
					Url:         "1url1",
				},
			},
			want:    &proto.NewShopResponse{},
			wantErr: false,
		},
		{
			name: "name is required",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.NewShopRequest{
					Name:        "",
					Description: "description",
					Url:         "url",
				},
			},
			want:    &proto.NewShopResponse{},
			wantErr: true,
		},
		{
			name: "shop error",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"NewShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(errors.New("shop error"))
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.NewShopRequest{
					Name:        "name",
					Description: "description",
					Url:         "url",
				},
			},
			want:    &proto.NewShopResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().NewShop(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestServerAPI_AlterShop(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.AlterShopRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.AlterShopResponse
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.AlterShopRequest{
					ShopId: 1,
					ShopInfo: &proto.ShopInfo{
						ShopId:      1,
						Name:        "name",
						Description: "description",
						Url:         "url",
					},
				},
			},
			want:    &proto.AlterShopResponse{},
			wantErr: false,
		},
		{
			name: "shop id is required",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.AlterShopRequest{
					ShopId: 0,
					ShopInfo: &proto.ShopInfo{
						ShopId:      0,
						Name:        "name",
						Description: "description",
						Url:         "url",
					},
				},
			},
			want:    &proto.AlterShopResponse{},
			wantErr: true,
		},
		{
			name: "shop error",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"AlterShop",
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
					mock.Anything,
				).Return(errors.New("shop error"))
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.AlterShopRequest{
					ShopId: 1,
					ShopInfo: &proto.ShopInfo{
						ShopId:      1,
						Name:        "name",
						Description: "description",
						Url:         "url",
					},
				},
			},
			want:    &proto.AlterShopResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().AlterShop(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestServerAPI_DeleteShop(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.DeleteShopRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.DeleteShopResponse
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"DeleteShop",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.DeleteShopRequest{
					ShopId: 1,
				},
			},
			want:    &proto.DeleteShopResponse{},
			wantErr: false,
		},
		{
			name: "shop id is required",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"DeleteShop",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.DeleteShopRequest{
					ShopId: 0,
				},
			},
			want:    &proto.DeleteShopResponse{},
			wantErr: true,
		},
		{
			name: "shop error",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"DeleteShop",
					mock.Anything,
					mock.Anything,
				).Return(errors.New("shop error"))
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.DeleteShopRequest{
					ShopId: 1,
				},
			},
			want:    &proto.DeleteShopResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().DeleteShop(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestServerAPI_ListShops(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.ListShopsRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.ListShopsResponse
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"ListShops",
					mock.Anything,
				).Return(nil, nil)
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in:  &proto.ListShopsRequest{},
			},
			want:    &proto.ListShopsResponse{},
			wantErr: false,
		},
		{
			name: "shop error",
			s: func() *service.ServerAPI {
				shop := &Shop{}
				shop.Mock.On(
					"ListShops",
					mock.Anything,
				).Return(nil, errors.New("shop error"))
				return &service.ServerAPI{Shop: shop}
			},
			args: args{
				ctx: context.TODO(),
				in:  &proto.ListShopsRequest{},
			},
			want:    &proto.ListShopsResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s().ListShops(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
