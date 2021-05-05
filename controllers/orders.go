package controllers

import (
	"api-orders/forms"
	"api-orders/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrdersController struct{}

var orderModel = new(services.OrderService)

//
// @ID api-orders-create-order
// @Summary Create an order
// @Description Create an order with status Created
// @Tags orders
// @Accept  json
// @Produce  json
// @Param user body forms.CreateOrder true "Add Product"
// @Success 200 {string} string	"id"
// @Router /orders [post]
// @Security ApiKeyAuth
func (controller OrdersController) Create(c *gin.Context) {
	user, _ := c.MustGet("User").(services.User)

	var input forms.CreateOrder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := orderModel.Create(input, user.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
	return
}

// @Summary Find an order
// @ID api-orders-read-order
// @Description Get an created order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Order
// @Router /orders/{id} [get]
// @Security ApiKeyAuth
func (controller OrdersController) FindOne(c *gin.Context) {
	user, _ := c.MustGet("User").(services.User)

	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	orderId, _ := primitive.ObjectIDFromHex(id)
	order, err := orderModel.FindOneWithUserId(orderId, user.UserId)

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
// @ID api-orders-cancel-order
// @Description Get an created order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Success 200 {object} models.Order
// @Router /orders/{id}/cancel [post]
// @Security ApiKeyAuth
func (controller OrdersController) Cancel(c *gin.Context) {
	user, _ := c.MustGet("User").(services.User)

	id := c.Param("id")
	if !primitive.IsValidObjectID(id) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	orderId, _ := primitive.ObjectIDFromHex(id)

	order, err := orderModel.Cancel(orderId, user.UserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": order})
	return
}
