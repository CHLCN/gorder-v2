package app

import "github.com/CHLCN/gorder-v2/order/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct{}

type Queries struct {
	GetCustormerOrder query.GetCustomerOrderHandler
}
