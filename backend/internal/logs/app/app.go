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

		var logs []*proto.Log
		err := json.Unmarshal([]byte(info.Info), &logs)

		if err != nil {
			return nil, fmt.Errorf("%s: %v", "failed unmarshalling logs: ", err)
		}
		for _, log := range logs {
			res = append(res, log)
		}

	}
	return &proto.GetLogsResponse{Logs: res}, nil
}

func (logs *Logs) GetGraph(
	ctx context.Context,
	content *proto.GetGraphRequest,
) (*proto.GetGraphResponse, error) {
	dateFrom := content.GetDateFrom()
	dateTo := content.GetDateTo()
	var res proto.GetGraphResponse
	var oneIntervalLater time.Time

	err := logs.logsStorage.GetMinDate(ctx, dateFrom)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if dateTo > time.Now().Unix() {
		return nil, fmt.Errorf("%s", "dateTo can't be more than current time ")
	}
	for dateFrom < dateTo {
		if content.GetInterval() == 0 {
			oneIntervalLater = time.Unix(dateFrom, 0).AddDate(0, 0, 1)
		} else if content.GetInterval() == 1 {
			oneIntervalLater = time.Unix(dateFrom, 0).AddDate(0, 1, 0)
		} else {
			return nil, fmt.Errorf("failed to get interval: invalid interval")
		}

		sales, err := logs.logsStorage.GetSales(ctx, dateFrom, oneIntervalLater.Unix())
		if err != nil {
			return nil, fmt.Errorf("%s: %v", "failed to get sales: ", err)
		}
		dateFrom = oneIntervalLater.Unix()
		var temp proto.Graph
		var flag bool = false
		for _, sale := range sales {
			if !flag {
				temp.Date = sale.Date
				flag = true
			}
			temp.TotalSales += sale.TotalSales
			temp.TotalQuantity += sale.TotalQuantity
		}
		res.Graphs = append(res.Graphs, &temp)

	}
	return &res, nil
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
