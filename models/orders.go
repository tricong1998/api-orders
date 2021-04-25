package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type OrderStatus int

type Product struct {
	name   string
	amount int
	price  float32
}

const (
	CREATED OrderStatus = iota
	CONFIRMED
	CANCELLED
	DELIVERED
)

type Order struct {
	id       primitive.ObjectID
	status   OrderStatus
	products []Product
}
