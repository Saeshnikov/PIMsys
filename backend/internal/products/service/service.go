package shop_service

import (
	"context"
	"fmt"
	proto "pim-sys/gen/go/products"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	proto.UnimplementedProductServer // Хитрая штука, о ней ниже
	product                          Product
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Product interface {
	NewProduct(
		ctx context.Context,
		content *proto.ProductInfo,
	) error
	AlterProduct(
		ctx context.Context,
		content *proto.ProductInfoWithId,
	) error
	DeleteProduct(
		ctx context.Context,
		content *proto.DeleteProductRequest,
	) error
	ListProducts(
		ctx context.Context,
	) (
		[]*proto.ProductInfoWithId,
		error,
	)
	SellProduct(
		ctx context.Context,
		content *proto.SellProductRequest,
	) error
}

func Register(gRPCServer *grpc.Server, product Product) {
	proto.RegisterProductServer(gRPCServer, &ServerAPI{product: product})
}

func (s *ServerAPI) NewProduct(
	ctx context.Context,
	in *proto.ProductInfo,
) (*proto.Empty, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	if len(in.Attributes) == 0 {
		return nil, status.Error(codes.InvalidArgument, "attributes are required")
	}

	if in.BranchId == 0 {
		return nil, status.Error(codes.InvalidArgument, "branch id is required")
	}

	if in.CategoryId == 0 {
		return nil, status.Error(codes.InvalidArgument, "code id is required")
	}

	if in.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "price is required")
	}
	if in.Amount < 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be positive or null")
	}

	err := s.product.NewProduct(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create new product: %w", err).Error())
	}

	return &proto.Empty{}, nil
}

func (s *ServerAPI) AlterProduct(
	ctx context.Context,
	in *proto.ProductInfoWithId,
) (*proto.Empty, error) {
	if in.ProductId == 0 {
		return nil, status.Error(codes.InvalidArgument, "product id is required")
	}

	if in.Product.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "price is required")
	}

	if in.Product.Amount < 0 {
		return nil, status.Error(codes.InvalidArgument, "amount must be positive or null")
	}

	err := s.product.AlterProduct(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to alter product: %w", err).Error())
	}

	return &proto.Empty{}, nil
}

func (s *ServerAPI) DeleteProduct(
	ctx context.Context,
	in *proto.DeleteProductRequest,
) (*proto.Empty, error) {
	if in.ProductId == 0 {
		return nil, status.Error(codes.InvalidArgument, "product id is required")
	}

	err := s.product.DeleteProduct(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete product: %w", err).Error())
	}

	return &proto.Empty{}, nil
}

func (s *ServerAPI) ListProducts(
	ctx context.Context,
	in *proto.Empty,
) (*proto.Products, error) {
	products, err := s.product.ListProducts(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to list product: %w", err).Error())
	}

	return &proto.Products{Product: products}, nil
}

func (s *ServerAPI) SellProduct(
	ctx context.Context,
	in *proto.SellProductRequest,
) (*proto.Empty, error) {
	if in.ProductId == 0 {
		return nil, status.Error(codes.InvalidArgument, "product id is required")
	}

	if in.Amount == 0 {
		return nil, status.Error(codes.InvalidArgument, "amount is required")
	}

	err := s.product.SellProduct(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to sell product: %w", err).Error())
	}

	return &proto.Empty{}, nil
}
