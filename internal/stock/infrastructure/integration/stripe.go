package integration

import (
	"context"

	_ "github.com/CHLCN/gorder-v2/common/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/stripe/stripe-go/v80"
	"github.com/stripe/stripe-go/v80/product"
)

type StripeAPI struct {
	apiKey string
}

func NewStripeAPI() *StripeAPI {
	key := viper.GetString("stripe-key")
	if key == "" {
		logrus.Fatal("empty key")
	}
	return &StripeAPI{apiKey: key}
}

func (s *StripeAPI) GetPriceByProductID(ctx context.Context, pid string) (string, error) {
	stripe.Key = s.apiKey
	result, err := product.Get(pid, &stripe.ProductParams{})
	if err != nil {
		return "", err
	}
	return result.DefaultPrice.ID, err
}

func (s *StripeAPI) GetProductByID(ctx context.Context, pid string) (*stripe.Product, error) {
	stripe.Key = s.apiKey
	return product.Get(pid, &stripe.ProductParams{})
}
