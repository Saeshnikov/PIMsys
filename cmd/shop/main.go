package main

import (
	"os"
	"os/signal"
	"pim-sys/internal/config"
	"pim-sys/internal/logger"
	shop_app "pim-sys/internal/shop/app"
	"syscall"
)

func main() {
	conf, err := config.InitConfig("configs/shop_config.yaml")
	if err != nil {
		panic(err)
	}

	log := logger.SetupLogger(conf.LogLevel)

	application := shop_app.New(log, conf.Grpc.Port, conf.TokenTTL)

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
