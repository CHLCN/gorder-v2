package service

import (
	"context"

	"github.com/CHLCN/gorder-v2/common/metrics"
	"github.com/CHLCN/gorder-v2/stock/adapters"
	"github.com/CHLCN/gorder-v2/stock/app"
	"github.com/CHLCN/gorder-v2/stock/app/query"
	"github.com/CHLCN/gorder-v2/stock/infrastructure/integration"
	"github.com/CHLCN/gorder-v2/stock/infrastructure/persistent"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(_ context.Context) app.Application {
	//stockRepo := adapters.NewMemoryStockRepository()
	db := persistent.NewMySQL()
	stockRepo := adapters.NewMySQLStockRepository(db)
	stripeAPI := integration.NewStripeAPI()
	metricsClient := metrics.NewPrometheusMetricsClient(&metrics.PrometheusMetricsClientConfig{
		Host:        viper.GetString("stock.metrics_export_addr"),
		ServiceName: viper.GetString("stock.service-name"),
	})
	return app.Application{
		Commands: app.Commands{},
		Queries: app.Queries{
			CheckIfItemsInStock: query.NewCheckIfItemsInStockHandler(stockRepo, stripeAPI, logrus.StandardLogger(), metricsClient),
			GetItems:            query.NewGetItemsHandler(stockRepo, logrus.StandardLogger(), metricsClient),
		},
	}
}
