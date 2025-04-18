package query

import (
	"context"
	"github.com/CHLCN/gorder-v2/common/decorator"
	domain "github.com/CHLCN/gorder-v2/stock/domain/stock"
	"github.com/CHLCN/gorder-v2/stock/entity"
	"github.com/sirupsen/logrus"
)

type GetItems struct {
	ItemIDs []string
}

type GetItemsHandler decorator.QueryHandler[GetItems, []*entity.Item]

type getItemsHandler struct {
	stockRepo domain.Repository
}

func NewGetItemsHandler(
	stockRepo domain.Repository,
	logger *logrus.Entry,
	metricsClient decorator.MetricsClient,
) GetItemsHandler {
	if stockRepo == nil {
		panic("nil stockRepo")
	}
	return decorator.ApplyQeuryDecorators[GetItems, []*entity.Item](
		getItemsHandler{stockRepo: stockRepo},
		logger,
		metricsClient,
	)
}

func (g getItemsHandler) Handle(ctx context.Context, query GetItems) ([]*entity.Item, error) {
	items, err := g.stockRepo.GetItems(ctx, query.ItemIDs)
	if err != nil {
		return nil, err
	}
	return items, nil
}
