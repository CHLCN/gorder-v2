package service

import (
	"context"
	grpcClient "github.com/CHLCN/gorder-v2/common/client"
	"github.com/CHLCN/gorder-v2/common/metrics"
	"github.com/CHLCN/gorder-v2/order/adapters"
	"github.com/CHLCN/gorder-v2/order/adapters/grpc"
	"github.com/CHLCN/gorder-v2/order/app"
	"github.com/CHLCN/gorder-v2/order/app/command"
	"github.com/CHLCN/gorder-v2/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) (app.Application, func()) {

	stockClient, closeStockClient, err := grpcClient.NewStockGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	stockGRPC := grpc.NewStockGRPC(stockClient)
	return newApplication(ctx, stockGRPC), func() {
		_ = closeStockClient()
	}

}

func newApplication(ctx context.Context, stockGRPC query.StockService) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{
			CreateOrder: command.NewCreateOrderHandler(orderRepo, stockGRPC, logger, metricClient),
			UpdateOrder: command.NewUpdateOrderHandler(orderRepo, logger, metricClient),
		},
		Queries: app.Queries{
			GetCustormerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient),
		},
	}
}
