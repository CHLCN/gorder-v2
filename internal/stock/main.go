package main

import (
	"context"
	"github.com/CHLCN/gorder-v2/common/tracing"

	_ "github.com/CHLCN/gorder-v2/common/config"
	"github.com/CHLCN/gorder-v2/common/discovery"
	"github.com/CHLCN/gorder-v2/common/genproto/stockpb"
	"github.com/CHLCN/gorder-v2/common/logging"
	"github.com/CHLCN/gorder-v2/common/server"
	"github.com/CHLCN/gorder-v2/stock/ports"
	"github.com/CHLCN/gorder-v2/stock/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	logging.Init()
}

func main() {
	serviceName := viper.GetString("stock.service-name")
	serverType := viper.GetString("stock.server-to-run")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdown, err := tracing.InitJaegerProvider(viper.GetString("jaeger.url"), serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer shutdown(ctx)

	application := service.NewApplication(ctx)

	deregisterFunc, err := discovery.RegisterToConsul(ctx, serviceName)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		_ = deregisterFunc()
	}()

	switch serverType {
	case "grpc":
		server.RunGRPCServer(serviceName, func(server *grpc.Server) {
			svc := ports.NewGRPCServer(application)
			stockpb.RegisterStockServiceServer(server, svc)
		})
	case "http":
		// TODO:
	default:
		panic("unexpected server type")
	}

}
