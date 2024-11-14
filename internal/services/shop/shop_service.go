package shop_service

import (
	"context"
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
	) (message string, shopId int32, err error)
	AlterShop(
		ctx context.Context,
		shopId int32,
		name string,
		description string,
		url string,
	) (message string, err error)
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

	if in.Description == "" {
		return nil, status.Error(codes.InvalidArgument, "description is required")
	}

	message, shopId, err := s.shop.NewShop(ctx, in.GetName(), in.GetDescription(), in.GetUrl())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create new shop")
	}

	return &proto.NewShopResponse{Message: message, ShopId: shopId}, nil
}

func (s *ServerAPI) AlterShop(
	ctx context.Context,
	in *proto.AlterShopRequest,
) (*proto.AlterShopResponse, error) {
	if in.ShopId == 0 {
		return nil, status.Error(codes.InvalidArgument, "shop id is required")
	}

	if in.ShopInfo.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "shop name is required")
	}

	message, err := s.shop.AlterShop(ctx, in.GetShopId(), in.ShopInfo.GetName(), in.ShopInfo.GetDescription(), in.ShopInfo.Url)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to register user")
	}

	return &proto.AlterShopResponse{Message: message}, nil
}
