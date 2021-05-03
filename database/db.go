package database

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatastore struct {
	Db      *mongo.Database
	Session *mongo.Client
}

const (
	// Timeout operations after N seconds
	connectTimeout           = 5
	connectionStringTemplate = "mongodb://%s:%s@%s"
)

func NewDatastore(databaseName string) *MongoDatastore {

	var mongoDataStore *MongoDatastore
	db, session := connect(databaseName)
	if db != nil && session != nil {

		// log statements here as well

		mongoDataStore = new(MongoDatastore)
		mongoDataStore.Db = db
		mongoDataStore.Session = session
		return mongoDataStore
	}

	return nil
}

func connect(databaseName string) (a *mongo.Database, b *mongo.Client) {
	var connectOnce sync.Once
	var db *mongo.Database
	var session *mongo.Client
	connectOnce.Do(func() {
		db, session = getConnection(databaseName)
	})

	return db, session
}

func getConnection(databaseName string) (*mongo.Database, *mongo.Client) {
	clusterEndpoint := "mongodb://127.0.0.1:27017"

	//os.Getenv("MONGODB_URI")

	connectionURI := clusterEndpoint

	//fmt.Sprintf(connectionStringTemplate, clusterEndpoint)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Printf("Failed to create client: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), connectTimeout*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}

	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}

	db := client.Database(databaseName)

	fmt.Println("Connected to MongoDB!")
	return db, client
}
