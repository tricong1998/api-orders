package services

import (
	"api-orders/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type PaymentServices struct{}

func (paymentServices PaymentServices) SendOrderToPayment(order models.Order) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	paymentUrl := os.Getenv("API_PAYMENTS_HOST") + os.Getenv("API_PAYMENTS_PORT") + "/api-payments/transactions"

	postBody := new(bytes.Buffer)
	json.NewEncoder(postBody).Encode(order)
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(paymentUrl, "application/json", postBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
		return err
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	sb := string(body)
	log.Printf(sb)

	return nil
}
