package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	proto "pim-sys/gen/go/shop"
	"pim-sys/internal/config"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T                  // Потребуется для вызова методов *testing.T
	Conf       *config.Config   // Конфигурация приложения
	ShopClient proto.ShopClient // Клиент для взаимодействия с gRPC-сервером Auth
}

const configPath = "suite/config.yaml"

const (
	grpcHost = "shop"
)

// New creates new test suite.
func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()   // Функция будет восприниматься как вспомогательная для тестов
	t.Parallel() // Разрешаем параллельный запуск тестов

	// Читаем конфиг из файла
	conf, err := config.InitConfig(configPath)
	require.NoError(t, err)

	// Основной родительский контекст
	ctx, cancelCtx := context.WithTimeout(context.Background(), conf.Grpc.Timeout)

	// Когда тесты пройдут, закрываем контекст
	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	// Адрес нашего gRPC-сервера
	grpcAddress := net.JoinHostPort(grpcHost, strconv.Itoa(conf.Grpc.Port))

	// Создаем клиент
	cc, err := grpc.DialContext(context.Background(),
		grpcAddress,
		// Используем insecure-коннект для тестов
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	// gRPC-клиент сервера Auth
	shopClient := proto.NewShopClient(cc)

	return ctx, &Suite{
		T:          t,
		Conf:       conf,
		ShopClient: shopClient,
	}
}