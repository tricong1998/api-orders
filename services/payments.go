package services

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type PaymentServices struct{}

type Product struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

type OrderPayment struct {
	Id              string    `json:"id"`
	Status          string    `json:"status"`
	Products        []Product `json:"products"`
	UserId          string    `json:"userId"`
	IsSendToPayment bool      `json:"isSendToPayment"`
}

func (paymentServices PaymentServices) SendOrderToPayment(order OrderPayment) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	paymentUrl := os.Getenv("API_PAYMENTS_HOST") + os.Getenv("API_PAYMENTS_PORT")
	client := &http.Client{
		// Set timeout to abort if the request takes too long
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("POST", paymentUrl, order)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err})
	}

	// Make website request call
	resp, err := client.Do(request)

	// If we have a successful request
	if resp.StatusCode == 200 {

	}
}
