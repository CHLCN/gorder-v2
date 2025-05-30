package service

import (
	"context"

	grpcClient "github.com/CHLCN/gorder-v2/common/client"
	"github.com/CHLCN/gorder-v2/common/metrics"
	"github.com/CHLCN/gorder-v2/payment/adapters"
	"github.com/CHLCN/gorder-v2/payment/app"
	"github.com/CHLCN/gorder-v2/payment/app/command"
	"github.com/CHLCN/gorder-v2/payment/domain"
	"github.com/CHLCN/gorder-v2/payment/infrastructure/processor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	orderClient, closeOrderClient, err := grpcClient.NewOrderGRPCClient(ctx)
	if err != nil {
		panic(err)
	}
	orderGRPC := adapters.NewOrderGRPC(orderClient)
	//memoryProcessor := processor.NewInmemProcessor()
	stripeProcessor := processor.NewStripeProcessor(viper.GetString("stripe-key"))
	return newApplication(ctx, orderGRPC, stripeProcessor), func() {
		_ = closeOrderClient()
	}
}

func newApplication(_ context.Context, orderGRPC command.OrderService, processor domain.Processor) app.Application {
	metricClient := metrics.NewPrometheusMetricsClient(&metrics.PrometheusMetricsClientConfig{
		Host:        viper.GetString("payment.metrics_export_addr"),
		ServiceName: viper.GetString("payment.service-name"),
	})
	return app.Application{
		Commands: app.Commands{
			CreatePayment: command.NewCreatePaymentHandler(processor, orderGRPC, logrus.StandardLogger(), metricClient),
		},
	}
}
