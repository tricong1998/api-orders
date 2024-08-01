package services

import (
	"api-orders/database"
	"os"
)

// Mongo server ip -> localhost -> 127.0.0.1 -> 0.0.0.0
var server = os.Getenv("DATABASE")

// Database name
var databaseName = "api-orders"

// Create a connection
var dbConnect = database.NewDatastore(databaseName)
