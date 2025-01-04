package shop_app

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	proto "pim-sys/gen/go/logs"
	auth_interceptor "pim-sys/internal/auth-interceptor"
	grpcapp "pim-sys/internal/grpc"
	shop_service "pim-sys/internal/logs/service"
	"pim-sys/internal/logs/storage"

	"google.golang.org/grpc"
)

type App struct {
	GRPCServer *grpcapp.App
}

type Logs struct {
	logsStorage *storage.Storage
}

func (logs *Logs) GetLogs(
	ctx context.Context,
	content *proto.GetLogsRequest,
) (*proto.GetLogsResponse, error) {

	user_id, err := auth_interceptor.GetFromContext(ctx, "user_id")
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "getting user_id from context: ", err)
	}

	userId, err := strconv.Atoi(user_id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "converting uid to int: ", err)
	}

	logsInfo, err := logs.logsStorage.GetLogs(ctx, int32(userId))
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "failed getting logs: ", err)
	}
	var res []*proto.Log
	for _, info := range logsInfo {
		var log *proto.Log
		err := json.Unmarshal([]byte(info.Info), log)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", "failed unmarshalling logs: ", err)
		}
		res = append(res, log)
	}
	return &proto.GetLogsResponse{Logs: res}, nil
}

func (logs *Logs) GetGraph(
	ctx context.Context,
	content *proto.GetGraphRequest,
) (*proto.GetGraphResponse, error) {
	logsInfo, err := logs.GetLogs(
		ctx,
		&proto.GetLogsRequest{
			ProductId: content.GetProductId(),
			DateFrom:  content.GetDateFrom(),
			DateTo:    content.GetDateFrom(),
		},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", "failed get logs: ", err)
	}
	switch content.GetInterval() {
	case 0:
	}
}

func New(
	log *slog.Logger,
	grpcPort int,
	connectionString string,
	tokenTTL time.Duration,
) *App {

	logsStorage, err := storage.New(connectionString)
	if err != nil {
		panic(err)
	}

	registerShop := func(gRPCServer *grpc.Server) {
		shop_service.Register(
			gRPCServer,
			&Logs{
				logsStorage: logsStorage,
			},
		)
	}

	grpcApp := grpcapp.New(log, registerShop, grpcPort, auth_interceptor.AuthInterceptor())

	return &App{
		GRPCServer: grpcApp,
	}
}
