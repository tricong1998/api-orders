package services

import (
	"api-orders/forms"
	"api-orders/models"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OrderService struct{}

func (model OrderService) Create(input forms.CreateOrder, userId string) (primitive.ObjectID, error) {
	var products []models.Product
	for _, v := range input.Products {
		products = append(products, models.Product{Name: v.Name, Amount: v.Amount, Price: v.Price})
	}
	order := models.Order{
		Id:       primitive.NewObjectID(),
		Status:   models.ORDER_STATUS_CREATED,
		Products: products,
		UserId:   userId,
	}

	collection := getCollection()
	result, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		log.Printf("Could not create Order: %v", err)
		return primitive.NilObjectID, err
	}
	go model.SendOrderToPayment(result.InsertedID.(primitive.ObjectID))
	return result.InsertedID.(primitive.ObjectID), err
}

func (model OrderService) SendOrderToPayment(orderId primitive.ObjectID) {
	var order models.Order
	collection := getCollection()

	result := collection.FindOne(context.TODO(), bson.D{bson.E{"_id", orderId}})
	result.Decode(&order)

	paymentService := new(PaymentServices)

	err := paymentService.SendOrderToPayment(order)

	if err != nil {
		log.Fatalln(err)
		return
	}

	collection.UpdateByID(context.TODO(), orderId, bson.D{bson.E{"status", models.ORDER_STATUS_CONFIRMED}})
}

func (model OrderService) FindOneById(id primitive.ObjectID) (*models.Order, error) {
	var order models.Order

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

func (model OrderService) FindOneWithUserId(id primitive.ObjectID, userId string) (*models.Order, error) {
	var order models.Order

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

func (model OrderService) CancelById(id primitive.ObjectID) (*models.Order, error) {
	var order models.Order

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
	if order.Status != models.ORDER_STATUS_DELIVERED {
		return nil, errors.New("Could not cancel Order if Order status is not CREATED")
	}
	var updatedDocument models.Order
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", models.ORDER_STATUS_CANCELLED}}}}
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

func (model OrderService) Cancel(id primitive.ObjectID, userId string) (*models.Order, error) {
	var order models.Order

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
	if order.Status != models.ORDER_STATUS_CREATED {
		return nil, errors.New("Could not cancel Order if Order status is not CREATED")
	}
	var updatedDocument models.Order
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"status", models.ORDER_STATUS_DELIVERED}}}}
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
	return dbConnect.Db.Collection(models.ORDER_COLLECTION_NAME)
}
