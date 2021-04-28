package models

import (
	"api-orders/forms"
)

type OrderStatus int

type Product struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

const (
	CREATED OrderStatus = iota
	CONFIRMED
	CANCELLED
	DELIVERED
)

type Order struct {
	Id       string      `json:"id"`
	Status   OrderStatus `json:"status"`
	Products []Product   `json:"products"`
}

func (model Order) Create(input forms.CreateOrder) Order {
	var products []Product
	for _, v := range input.Products {
		products = append(products, Product{Name: v.Name, Amount: v.Amount, Price: v.Price})
	}
	order := Order{
		Id:       "id",
		Status:   CREATED,
		Products: products,
	}
	return order
}
