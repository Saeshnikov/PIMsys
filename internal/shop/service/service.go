package shop_service

import (
	"context"
	"fmt"
	proto "pim-sys/gen/go/shop"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	proto.UnimplementedShopServer // Хитрая штука, о ней ниже
	shop                          Shop
}

// Тот самый интерфейс, котрый мы передавали в grpcApp
type Shop interface {
	NewShop(
		ctx context.Context,
		name string,
		description string,
		url string,
	) error
	AlterShop(
		ctx context.Context,
		shopId int32,
		name string,
		description string,
		url string,
	) error
	DeleteShop(
		ctx context.Context,
		shopId int32,
	) error
	ListShops(
		ctx context.Context,
	) (
		[]*proto.ShopInfo,
		error,
	)
}

func Register(gRPCServer *grpc.Server, shop Shop) {
	proto.RegisterShopServer(gRPCServer, &ServerAPI{shop: shop})
}

func (s *ServerAPI) NewShop(
	ctx context.Context,
	in *proto.NewShopRequest,
) (*proto.NewShopResponse, error) {
	if in.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "name is required")
	}

	err := s.shop.NewShop(ctx, in.GetName(), in.GetDescription(), in.GetUrl())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to create new shop: %w", err).Error())
	}

	return &proto.NewShopResponse{}, nil
}

func (s *ServerAPI) AlterShop(
	ctx context.Context,
	in *proto.AlterShopRequest,
) (*proto.AlterShopResponse, error) {
	if in.ShopId == 0 {
		return nil, status.Error(codes.InvalidArgument, "shop id is required")
	}

	err := s.shop.AlterShop(ctx, in.GetShopId(), in.ShopInfo.GetName(), in.ShopInfo.GetDescription(), in.ShopInfo.Url)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to alter shop: %w", err).Error())
	}

	return &proto.AlterShopResponse{}, nil
}

func (s *ServerAPI) DeleteShop(
	ctx context.Context,
	in *proto.DeleteShopRequest,
) (*proto.DeleteShopResponse, error) {
	if in.ShopId == 0 {
		return nil, status.Error(codes.InvalidArgument, "shop id is required")
	}

	err := s.shop.DeleteShop(ctx, in.GetShopId())
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete shop: %w", err).Error())
	}

	return &proto.DeleteShopResponse{}, nil
}

func (s *ServerAPI) ListShops(
	ctx context.Context,
	in *proto.ListShopsRequest,
) (*proto.ListShopsResponse, error) {
	shopInfo, err := s.shop.ListShops(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to delete shop: %w", err).Error())
	}

	return &proto.ListShopsResponse{Info: shopInfo}, nil
}
