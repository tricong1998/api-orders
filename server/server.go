package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitServer() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	route := initRoute()
	fmt.Println(os.Getenv("PORT"))
	route.Run(":" + os.Getenv("PORT"))
}
