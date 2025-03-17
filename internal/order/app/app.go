package app

import (
	"github.com/CHLCN/gorder-v2/order/app/command"
	"github.com/CHLCN/gorder-v2/order/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateOrder command.CreateOrderHandler
	UpdateOrder command.UpdateOrderHandler
}

type Queries struct {
	GetCustormerOrder query.GetCustomerOrderHandler
}
