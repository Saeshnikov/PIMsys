package product_app_test

import (
	"context"
	"log/slog"
	proto "pim-sys/gen/go/products"
	app "pim-sys/internal/products/app"
	"testing"
	"time"

	mock "github.com/stretchr/testify/mock"
	"google.golang.org/grpc/metadata"
)

//go:generate mockery --testonly --name Storage --dir . --outpkg product_app_test --output .

func TestProducts_NewProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.ProductInfo
	}
	tests := []struct {
		name     string
		products func() *app.Products
		args     args
		wantErr  bool
	}{
		{
			name: "deafult",
			products: func() *app.Products {
				storage := &Storage{}
				storage.Mock.On(
					"GetAccessableBranchIds",
					mock.Anything,
					mock.Anything).Return([]int32{1}, nil)
				storage.Mock.On(
					"CreateProduct",
					mock.Anything,
					mock.Anything).Return(int32(1), nil)
				storage.Mock.On(
					"AlterAttributes",
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)

				return &app.Products{
					ProductsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.ProductInfo{BranchId: 1, Attributes: []*proto.Attribute{{}}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.products().NewProduct(tt.args.ctx, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Products.NewProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducts_AlterProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.ProductInfoWithId
	}
	tests := []struct {
		name     string
		products func() *app.Products
		args     args
		wantErr  bool
	}{
		{
			name: "deafult",
			products: func() *app.Products {
				storage := &Storage{}
				storage.Mock.On(
					"ListProducts",
					mock.Anything,
					mock.Anything).Return([]*proto.ProductInfoWithId{{ProductId: 1}}, nil)
				storage.Mock.On(
					"AlterProduct",
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)

				return &app.Products{
					ProductsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.ProductInfoWithId{ProductId: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.products().AlterProduct(tt.args.ctx, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Products.AlterProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducts_DeleteProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.DeleteProductRequest
	}
	tests := []struct {
		name     string
		products func() *app.Products
		args     args
		wantErr  bool
	}{
		{
			name: "deafult",
			products: func() *app.Products {
				storage := &Storage{}
				storage.Mock.On(
					"ListProducts",
					mock.Anything,
					mock.Anything).Return([]*proto.ProductInfoWithId{{ProductId: 1}}, nil)
				storage.Mock.On(
					"DeleteProduct",
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)

				return &app.Products{
					ProductsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.DeleteProductRequest{ProductId: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.products().DeleteProduct(tt.args.ctx, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Products.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProducts_SellProduct(t *testing.T) {
	type args struct {
		ctx     context.Context
		content *proto.SellProductRequest
	}
	tests := []struct {
		name     string
		products func() *app.Products
		args     args
		wantErr  bool
	}{
		{
			name: "deafult",
			products: func() *app.Products {
				storage := &Storage{}
				storage.Mock.On(
					"ListProducts",
					mock.Anything,
					mock.Anything).Return([]*proto.ProductInfoWithId{{ProductId: 1}}, nil)
				storage.Mock.On(
					"SellProduct",
					mock.Anything,
					mock.Anything,
					mock.Anything).Return(nil)

				return &app.Products{
					ProductsStorage: storage,
				}
			},
			args: args{
				ctx:     metadata.NewIncomingContext(context.TODO(), metadata.New(map[string]string{"user_id": "1"})),
				content: &proto.SellProductRequest{ProductId: 1},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.products().SellProduct(tt.args.ctx, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Products.SellProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		log              *slog.Logger
		grpcPort         int
		connectionString string
		tokenTTL         time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				log:              slog.Default(),
				grpcPort:         1111,
				connectionString: "str",
				tokenTTL:         time.Duration(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.New(tt.args.log, tt.args.grpcPort, tt.args.connectionString, tt.args.tokenTTL)
		})
	}
}
