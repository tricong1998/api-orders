package models

import "api-orders/api-orders/forms"

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
	id       string
	status   OrderStatus
	products []Product
}

func (model Order) Create(input forms.CreateOrder) Order {
	order := Order{
		id:       "id",
		status:   CREATED,
		products: input.Products,
	}

	return order
}
