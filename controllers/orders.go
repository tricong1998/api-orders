package controllers

import (
	"api-orders/api-orders/forms"
	"api-orders/api-orders/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersController struct{}

var orderModel = new(models.Order)

func (controller OrdersController) Create(c *gin.Context) {
	var input forms.CreateOrder
	if err := c.ShouldBindJson(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := orderModel.Create(input)

	c.JSON(order)
	return
}
