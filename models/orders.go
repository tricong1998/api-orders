package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderStatus int

type Product struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

const ORDER_STATUS_CREATED = "CREATED"
const ORDER_STATUS_CONFIRMED = "CONFIRMED"
const ORDER_STATUS_CANCELLED = "CANCELLED"
const ORDER_STATUS_DELIVERED = "DELIVERED"

const ORDER_COLLECTION_NAME = "orders"

type Order struct {
	Id              primitive.ObjectID `json:"id"`
	Status          string             `json:"status"`
	Products        []Product          `json:"products"`
	UserId          string             `json:"userId"`
	IsSendToPayment bool               `json:"isSendToPayment"`
}
