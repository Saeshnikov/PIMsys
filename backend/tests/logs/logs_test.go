package tests

import (
	"testing"

	"pim-sys/tests/assistance"
	suite "pim-sys/tests/logs/suite"

	proto "pim-sys/gen/go/logs"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

const configPath = "suite/config.yaml"

func TestRegisterLogin_Login_HappyPath(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	newLog, err := st.LogsClient.GetLogs(ctx, &proto.GetLogsRequest{
		ProductId: 1,
	})
	require.NoError(t, err)

	require.NotNil(t, newLog)

	newSales, err := st.LogsClient.GetGraph(ctx, &proto.GetGraphRequest{
		Interval: 0,
		DateFrom: 1704074400, //01.01.2024 02.00.00
		DateTo:   1735704000, //01.01.2025 02.00.00
	})
	require.NoError(t, err)

	require.NotNil(t, newSales)

	require.Equal(t, 160, int(newSales.Graphs[0].TotalSales))
}

func TestLogUpdateProduct(t *testing.T) {
	token := assistance.GetTestToken(t)
	ctx, st := suite.New(t, configPath)
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)

	udpdateLog, err := st.LogsClient.GetLogs(ctx, &proto.GetLogsRequest{
		ProductId: 1,
	})
	require.NoError(t, err)
	require.Equal(t, len(udpdateLog.GetLogs()), 2)
}
