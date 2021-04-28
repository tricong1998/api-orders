package controllers

import (
	"api-orders/forms"
	"api-orders/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdersController struct{}

var orderModel = new(models.Order)

func (controller OrdersController) Create(c *gin.Context) {
	var input forms.CreateOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := orderModel.Create(input)

	c.JSON(http.StatusOK, gin.H{"data": order})
	return
}
