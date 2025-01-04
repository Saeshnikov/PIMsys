package main

import (
	"os"
	"os/signal"
	"pim-sys/internal/config"
	"pim-sys/internal/logger"
	products_app "pim-sys/internal/products/app"
	"syscall"
)

func main() {
	conf, err := config.InitConfig("configs/products_config.yaml")
	if err != nil {
		panic(err)
	}

	log := logger.SetupLogger(conf.LogLevel)

	application := products_app.New(log, conf.Grpc.Port, conf.ConnectionString, conf.TokenTTL)

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
