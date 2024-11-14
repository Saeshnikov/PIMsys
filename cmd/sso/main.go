package main

import (
	"os"
	"os/signal"
	auth_app "pim-sys/internal/auth/app"
	"pim-sys/internal/config"
	"pim-sys/internal/logger"
	"syscall"
)

func main() {
	cfg, err := config.InitConfig("configs/auth_config.yaml")
	if err != nil {
		panic(err)
	}

	log := logger.SetupLogger(cfg.LogLevel)

	application := auth_app.New(log, cfg.Grpc.Port, cfg.ConnectionString, cfg.TokenTTL)

	go func() {
		application.GRPCServer.MustRun()
	}()

	// Graceful shutdown

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	application.GRPCServer.Stop()
	log.Info("Gracefully stopped")
}
