package server

import (
	"api-orders/controllers"

	"github.com/gin-gonic/gin"
)

func initRoute() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	api := route.Group("api")
	{
		orders := api.Group("orders")
		{
			order := new(controllers.OrdersController)
			orders.POST("", order.Create)
		}
	}
	return route
}
