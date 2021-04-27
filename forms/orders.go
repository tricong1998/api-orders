package forms

type Product struct {
	Name   string  `json:"name" binding:"required"`
	Amount int     `json:"amount" binding:"required"`
	Price  float32 `json:"price" binding:"required"`
}

type CreateOrder struct {
	Products []Product `json:"product" binding:"required"`
}
