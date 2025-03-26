package logger_test

import (
	"log/slog"
	"pim-sys/internal/logger"
	"testing"
)

func TestSetupLogger(t *testing.T) {
	type args struct {
		env string
	}
	tests := []struct {
		name string
		args args
		want *slog.Logger
	}{
		{
			name: "default",
			args: args{
				env: "local",
			},
		},
		{
			name: "default",
			args: args{
				env: "dev",
			},
		},
		{
			name: "default",
			args: args{
				env: "prod",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger.SetupLogger(tt.args.env)
		})
	}
}
