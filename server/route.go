package server

import (
	"api-orders/api-orders/controllers"

	"github.com/gin-gonic/gin"
)

func initRoute() *gin.Engine {
	route := gin.New()
	route.use(gin.Logger)
	route.use(gin.Recovery)

	api := route.group("api")
	{
		orders := api.group("orders")
		{
			order := new(controllers.OrdersController)
			orders.POST("", order.Create)
		}
	}
	return route
}
