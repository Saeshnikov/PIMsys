package product_app

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/products"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	product_service "pim-sys/internal/products/service"
	"pim-sys/internal/products/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Storage interface {
	CreateProduct(
		ctx context.Context,
		content *proto.ProductInfo,
	) (int32, error)
	AlterAttributes(
		ctx context.Context,
		productId int32,
		attr *proto.Attribute,
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
		userId int32,
	) (
		[]*proto.ProductInfoWithId,
		error,
	)
	SellProduct(
		ctx context.Context,
		content *proto.SellProductRequest,
	) error
	GetAccessableBranchIds(
		ctx context.Context,
		userId int32,
	) (
		[]int32,
		error,
	)
}

type Products struct {
	ProductsStorage Storage
}

func (products *Products) NewProduct(
	ctx context.Context,
	content *proto.ProductInfo,
) error {

	err := products.userMustHaveAccessToBranch(ctx, content.GetBranchId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}
	productId, err := products.ProductsStorage.CreateProduct(ctx, content)
	if err != nil {
		return fmt.Errorf("%s: %v", "creating product: ", err)
	}

	for _, attr := range content.Attributes {
		err = products.ProductsStorage.AlterAttributes(ctx, productId, attr)
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

	return products.ProductsStorage.AlterProduct(ctx, content)
}

func (products *Products) DeleteProduct(
	ctx context.Context,
	content *proto.DeleteProductRequest,
) error {

	err := products.userMustHaveAccessToProduct(ctx, content.GetProductId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}

	return products.ProductsStorage.DeleteProduct(ctx, content)
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

	return products.ProductsStorage.ListProducts(ctx, int32(userId))
}

func (products *Products) SellProduct(
	ctx context.Context,
	content *proto.SellProductRequest,
) error {
	err := products.userMustHaveAccessToProduct(ctx, content.GetProductId())
	if err != nil {
		return fmt.Errorf("%s: %v", "checking user permissions: ", err)
	}

	return products.ProductsStorage.SellProduct(ctx, content)
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
		product_service.Register(
			gRPCServer,
			&Products{
				ProductsStorage: productsStorage,
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

	availableProducts, err := products.ProductsStorage.GetAccessableBranchIds(ctx, int32(userId))
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
