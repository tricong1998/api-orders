package models

type OrderStatus int

type Product struct {
	name  string
	price float32
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
