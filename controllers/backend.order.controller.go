package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BackendOrdersController struct{}

// @Summary Find an order
// @ID api-orders-backend-read-order
// @Description Get an created order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Order
// @Router /backend/orders/{id} [get]
// @Security ApiKeyAuth
func (controller BackendOrdersController) FindOne(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	orderId, _ := primitive.ObjectIDFromHex(id)
	order, err := orderModel.FindOneById(orderId)

	if err != nil {
		if order == nil {
			c.JSON(http.StatusNotFound, gin.H{"msg": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
	return
}

// @Summary Cancel an order
// @ID api-orders-backend-cancel-order
// @Description Get an created order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Order
// @Router /backend/orders/{id}/cancel [post]
// @Security ApiKeyAuth
func (controller BackendOrdersController) Cancel(c *gin.Context) {
	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	orderId, _ := primitive.ObjectIDFromHex(id)

	order, err := orderModel.CancelById(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
	return
}
