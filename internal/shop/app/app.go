package shop_app

import (
	"context"
	"log/slog"
	"time"

	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	shop_service "pim-sys/internal/shop/service"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type handlers struct {
	str string
	num int32
}

func (h handlers) NewShop(
	ctx context.Context,
	name string,
	description string,
	url string,
) (string, int32, error) {
	return h.str, h.num, nil
}

func (h handlers) AlterShop(
	ctx context.Context,
	shopId int32,
	name string,
	description string,
	url string,
) (string, error) {
	return h.str, nil
}

func RegisterShop(gRPCServer *grpc.Server) {
	shop_service.Register(gRPCServer, handlers{str: "alala", num: 101010})
}

func New(
	log *slog.Logger,
	grpcPort int,
	tokenTTL time.Duration,
) *App {

	grpcApp := grpcapp.New(log, RegisterShop, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}
