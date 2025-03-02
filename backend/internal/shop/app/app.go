package shop_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/shop"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	shop_service "pim-sys/internal/shop/service"
	"pim-sys/internal/shop/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Storage interface {
	CreateShop(ctx context.Context, userId int, name string, description string, url string) error
	AlterShop(ctx context.Context, shopId int32, name string, description string, url string) error
	DeleteShop(ctx context.Context, shopId int32) error
	ListShops(ctx context.Context, userId int32) (
		[]*proto.ShopInfo,
		error,
	)
}

type Shop struct {
	ShopStorage Storage
}

func (shop *Shop) NewShop(
	ctx context.Context,
	name string,
	description string,
	url string,
) error {

	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	return shop.ShopStorage.CreateShop(ctx, userId, name, description, url)
}

func (shop *Shop) AlterShop(
	ctx context.Context,
	shopId int32,
	name string,
	description string,
	url string,
) error {
	err := shop.userMustHaveAccess(ctx, shopId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return shop.ShopStorage.AlterShop(ctx, shopId, name, description, url)
}

func (shop *Shop) DeleteShop(
	ctx context.Context,
	shopId int32,
) error {

	err := shop.userMustHaveAccess(ctx, shopId)
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions", err)
	}

	return shop.ShopStorage.DeleteShop(ctx, shopId)
}

func (shop *Shop) ListShops(
	ctx context.Context,
) (
	[]*proto.ShopInfo,
	error,
) {
	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	return shop.ShopStorage.ListShops(ctx, int32(userId))
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	shopStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerShop := func(gRPCServer *grpc.Server) {
		shop_service.Register(
			gRPCServer,
			&Shop{
				ShopStorage: shopStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerShop, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (shop *Shop) userMustHaveAccess(ctx context.Context, shopId int32) error {
	availableShops, err := shop.ListShops(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user's available shops: ", err)
	}

	for _, available := range availableShops {
		if shopId == available.GetShopId() {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or shop does not exist")
}
