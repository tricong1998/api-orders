package models

import (
	"api-orders/forms"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderStatus int

type Product struct {
	Name   string  `json:"name"`
	Amount int     `json:"amount"`
	Price  float32 `json:"price"`
}

const ORDER_STATUS_CREATED = "CREATED"
const ORDER_STATUS_CONFIRMED = "CONFIRMED"
const ORDER_STATUS_CANCELLED = "CANCELLED"
const ORDER_STATUS_DELIVERED = "DELIVERED"

const ORDER_COLLECTION_NAME = "orders"

type Order struct {
	Id              primitive.ObjectID `json:"id"`
	Status          string             `json:"status"`
	Products        []Product          `json:"products"`
	UserId          string             `json:"userId"`
	IsSendToPayment bool               `json:"isSendToPayment"`
}

func (model Order) Create(input forms.CreateOrder, userId string) (primitive.ObjectID, error) {
	var products []Product
	for _, v := range input.Products {
		products = append(products, Product{Name: v.Name, Amount: v.Amount, Price: v.Price})
	}
	order := Order{
		Id:       primitive.NewObjectID(),
		Status:   ORDER_STATUS_CREATED,
		Products: products,
		UserId:   userId,
	}

	collection := getCollection()
	result, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Printf("Could not create Order: %v", err)
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), err
}

func (model Order) FindOneById(id primitive.ObjectID) (*Order, error) {
	var order Order

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}})

	if result == nil {
		return nil, errors.New("Could not find a Order")
	}
	err := result.Decode(&order)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Order")
		}
		return nil, err
	}

	return &order, nil
}

func (model Order) FindOneWithUserId(id primitive.ObjectID, userId string) (*Order, error) {
	var order Order

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}, bson.E{"userId", userId}})

	if result == nil {
		return nil, errors.New("Could not find a Order")
	}
	err := result.Decode(&order)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Order")
		}
		return nil, err
	}

	return &order, nil
}

func (model Order) CancelById(id primitive.ObjectID) (*Order, error) {
	var order Order

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}})
	if result == nil {
		return nil, errors.New("Could not find an Order")
	}

	err := result.Decode(&order)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Orders: %v", order)
	if order.Status != ORDER_STATUS_DELIVERED {
		return nil, errors.New("Could not cancel Order if Order status is not CREATED")
	}
	var updatedDocument Order
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", ORDER_STATUS_CANCELLED}}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	updateError := getCollection().FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedDocument)

	if updateError != nil {
		if updateError == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Order")
		}
		log.Fatal(err)
	}
	return &updatedDocument, nil
}

func (model Order) Cancel(id primitive.ObjectID, userId string) (*Order, error) {
	var order Order

	result := getCollection().FindOne(context.TODO(), bson.D{bson.E{"_id", id}, bson.E{"userId", userId}})
	if result == nil {
		return nil, errors.New("Could not find an Order")
	}

	err := result.Decode(&order)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("Orders: %v", order)
	if order.Status != ORDER_STATUS_CREATED {
		return nil, errors.New("Could not cancel Order if Order status is not CREATED")
	}
	var updatedDocument Order
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", ORDER_STATUS_DELIVERED}}}}
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	updateError := getCollection().FindOneAndUpdate(context.TODO(), filter, update, &opt).Decode(&updatedDocument)

	if updateError != nil {
		if updateError == mongo.ErrNoDocuments {
			return nil, errors.New("Could not find an Order")
		}
		log.Fatal(err)
	}
	return &updatedDocument, nil
}

func getCollection() *mongo.Collection {
	return dbConnect.Db.Collection(ORDER_COLLECTION_NAME)
}
