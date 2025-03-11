package main

import (
	"github.com/CHLCN/gorder-v2/order/app"
	"github.com/CHLCN/gorder-v2/order/app/query"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServer struct {
	app app.Application
}

func (H HTTPServer) PostCustomerCustomerIDOrders(c *gin.Context, customerID string) {

}

func (H HTTPServer) GetCustomerCustomerIDOrdersOrderID(c *gin.Context, customerID string, orderID string) {
	o, err := H.app.Queries.GetCustormerOrder.Handle(c, query.GetCustomerOrder{
		OrderID:    "fake-ID",
		CustomerID: "fake-customer-id",
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": o})
}
