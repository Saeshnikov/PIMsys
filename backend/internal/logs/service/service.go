package logs_service

import (
	"context"
	"fmt"
	proto "pim-sys/gen/go/logs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerAPI struct {
	proto.UnimplementedShopServer // Хитрая штука, о ней ниже
	logs                          Logs
}

// Тот самый интерфейс, который мы передавали в grpcApp
type Logs interface { //!!!!!!!!!!!!!!
	GetLogs(
		ctx context.Context,
		content *proto.GetLogsRequest,
	) ([]*proto.Log, error)
	GetGraph(
		ctx context.Context,
		content *proto.GetGraphRequest,
	) ([]*proto.Graph, error)
}

func Register(gRPCServer *grpc.Server, logs Logs) {
	proto.RegisterShopServer(gRPCServer, &ServerAPI{logs: logs})
}

func (s *ServerAPI) GetLogs(
	ctx context.Context,
	in *proto.GetLogsRequest,
) (*proto.GetLogsResponse, error) {
	if in.ProductId >= 0 {
		return nil, status.Error(codes.InvalidArgument, "product_id cannot be less than equal to zero")
	}
	if in.DateFrom >= in.DateTo {
		return nil, status.Error(codes.InvalidArgument, "inaccurate start date")
	}
	logs, err := s.logs.GetLogs(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to get logs: %w", err).Error())
	}
	return &proto.GetLogsResponse{Logs: logs}, nil
}

func (s *ServerAPI) GetGraph(
	ctx context.Context,
	in *proto.GetGraphRequest,
) (*proto.GetGraphResponse, error) {
	if in.ProductId >= 0 {
		return nil, status.Error(codes.InvalidArgument, "product_id cannot be less than equal to zero")
	}
	if in.DateFrom >= in.DateTo {
		return nil, status.Error(codes.InvalidArgument, "inaccurate start date")
	}
	graphs, err := s.logs.GetGraph(ctx, in)
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Errorf("failed to get graph: %w", err).Error())
	}
	return &proto.GetGraphResponse{Graphs: graphs}, nil
}
