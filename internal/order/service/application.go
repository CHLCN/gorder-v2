package service

import (
	"context"
	"github.com/CHLCN/gorder-v2/common/metrics"
	"github.com/CHLCN/gorder-v2/order/adapters"
	"github.com/CHLCN/gorder-v2/order/app"
	"github.com/CHLCN/gorder-v2/order/app/query"
	"github.com/sirupsen/logrus"
)

func NewApplication(ctx context.Context) app.Application {
	orderRepo := adapters.NewMemoryOrderRepository()
	logger := logrus.NewEntry(logrus.StandardLogger())
	metricClient := metrics.TodoMetrics{}
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			GetCustormerOrder: query.NewGetCustomerOrderHandler(orderRepo, logger, metricClient),
		},
	}
}
