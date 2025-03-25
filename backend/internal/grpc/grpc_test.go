package grpcapp

import (
	"log/slog"
	"testing"

	"google.golang.org/grpc"
)

func TestNew(t *testing.T) {
	type args struct {
		log          *slog.Logger
		register     func(gRPCServer *grpc.Server)
		port         int
		interceptors []grpc.UnaryServerInterceptor
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		{
			name: "default",
			args: args{
				log:          slog.Default(),
				register:     func(gRPCServer *grpc.Server) {},
				port:         1111,
				interceptors: []grpc.UnaryServerInterceptor{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.log, tt.args.register, tt.args.port, tt.args.interceptors...)
		})
	}
}
