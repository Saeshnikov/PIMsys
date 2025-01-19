package shop_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/products"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	shop_service "pim-sys/internal/products/service"
	"pim-sys/internal/products/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Products struct {
	productsStorage *storage.Storage
}

func (products *Products) NewProduct(
	ctx context.Context,
	content *proto.ProductInfo,
) error {

	err := products.userMustHaveAccessToBranch(ctx, content.GetBranchId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}
	productId, err := products.productsStorage.CreateProduct(ctx, content)
	if err != nil {
		return fmt.Errorf("%s: %v", "creating product: ", err)
	}

	for _, attr := range content.Attributes {
		err = products.productsStorage.AlterAttributes(ctx, productId, attr)
		if err != nil {
			errDelete := products.DeleteProduct(ctx, &proto.DeleteProductRequest{ProductId: productId})
			if errDelete != nil {
				return fmt.Errorf("%s: %v", "unexpected error: ", err)
			}
			return fmt.Errorf("%s: %v", "add attributes: ", err)
		}
	}
	return nil
}

func (products *Products) AlterProduct(
	ctx context.Context,
	content *proto.ProductInfoWithId,
) error {
	err := products.userMustHaveAccessToProduct(ctx, content.GetProductId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}

	return products.productsStorage.AlterProduct(ctx, content)
}

func (products *Products) DeleteProduct(
	ctx context.Context,
	content *proto.DeleteProductRequest,
) error {

	err := products.userMustHaveAccessToProduct(ctx, content.GetProductId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}

	return products.productsStorage.DeleteProduct(ctx, content)
}

func (products *Products) ListProducts(
	ctx context.Context,
) (
	[]*proto.ProductInfoWithId,
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

	return products.productsStorage.ListProducts(ctx, int32(userId))
}

func (products *Products) SellProduct(
	ctx context.Context,
	content *proto.SellProductRequest,
) error {
	err := products.userMustHaveAccessToProduct(ctx, content.GetProductId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}

	return products.productsStorage.SellProduct(ctx, content)
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	productsStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerProduct := func(gRPCServer *grpc.Server) {
		shop_service.Register(
			gRPCServer,
			&Products{
				productsStorage: productsStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerProduct, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}

func (products *Products) userMustHaveAccessToProduct(ctx context.Context, productId int32) error {
	availableProducts, err := products.ListProducts(ctx)
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user's available shops: ", err)
	}

	for _, available := range availableProducts {
		if productId == available.GetProductId() {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or product does not exist")
}

func (products *Products) userMustHaveAccessToBranch(ctx context.Context, brachId int32) error {
	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	availableProducts, err := products.productsStorage.GetAccessableBranchIds(ctx, int32(userId))
	if err != nil {
		return fmt.Errorf("%s: %v", "getting user's available shops: ", err)
	}

	for _, available := range availableProducts {
		if brachId == available {
			return nil
		}
	}

	return fmt.Errorf("%s", "access denied or branch does not exist")
}
