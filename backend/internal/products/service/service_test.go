package product_service_test

import (
	"context"
	proto "pim-sys/gen/go/products"
	service "pim-sys/internal/products/service"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --testonly --name Product --dir . --outpkg product_service_test --output .

func TestServerAPI_NewProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.ProductInfo
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"NewProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.ProductInfo{
					CategoryId: 1,
					BranchId:   1,
					Status:     "status",
					Name:       "name",
					Price:      1.0,
					Amount:     1,
					Attributes: []*proto.Attribute{
						{},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s().NewProduct(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_AlterProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.ProductInfoWithId
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		wantErr bool
	}{
		{
			name: "default (pair_1)",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"AlterProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.ProductInfoWithId{
					ProductId: 1,
					Product: &proto.ProductInfo{
						CategoryId: 1,
						BranchId:   1,
						Status:     "status",
						Name:       "name",
						Price:      1.0,
						Amount:     1,
						Attributes: []*proto.Attribute{
							{},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "pair_2",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"AlterProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.ProductInfoWithId{
					ProductId: 1,
					Product: &proto.ProductInfo{
						Price:  0,
						Amount: 0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "pair_3",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"AlterProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.ProductInfoWithId{
					ProductId: 0,
					Product: &proto.ProductInfo{
						Price:  1.0,
						Amount: 0,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "pair_4",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"AlterProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.ProductInfoWithId{
					ProductId: 0,
					Product: &proto.ProductInfo{
						Price:  0,
						Amount: 1,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s().AlterProduct(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_DeleteProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.DeleteProductRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"DeleteProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.DeleteProductRequest{
					ProductId: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s().DeleteProduct(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_ListProducts(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.Empty
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"ListProducts",
					mock.Anything,
				).Return([]*proto.ProductInfoWithId{}, nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in:  &proto.Empty{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s().ListProducts(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestServerAPI_SellProduct(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *proto.SellProductRequest
	}
	tests := []struct {
		name    string
		s       func() *service.ServerAPI
		args    args
		want    *proto.Empty
		wantErr bool
	}{
		{
			name: "default",
			s: func() *service.ServerAPI {
				product := &Product{}
				product.Mock.On(
					"SellProduct",
					mock.Anything,
					mock.Anything,
				).Return(nil)
				return &service.ServerAPI{Product: product}
			},
			args: args{
				ctx: context.TODO(),
				in: &proto.SellProductRequest{
					ProductId: 1,
					Amount:    1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.s().SellProduct(tt.args.ctx, tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
