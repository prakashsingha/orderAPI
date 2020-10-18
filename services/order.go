package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order structure
type Order struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	HotelID      primitive.ObjectID `json:"hotelId,omitempty" bson:"hotelId,omitempty"`
	HotelName    string             `json:"hotelName,omitempty" bson:"hotelName,omitempty"`
	CheckInDate  time.Time          `json:"checkInDate,omitempty" bson:"checkInDate,omitempty"`
	CheckOutDate time.Time          `json:"checkOutDate,omitempty" bson:"checkOutDate,omitempty"`
	Customer     *Customer          `json:"customer,omitempty" bson:"customer,omitempty"`
	Room         *Room              `json:"room,omitempty" bson:"room,omitempty"`
	TotalAmount  float32            `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	PaymentID    primitive.ObjectID `json:"paymentId,omitempty" bson:"paymentId,omitempty"`
}

// SearchParam is the structure for filter parameter to search orders
type SearchParam struct {
	HotelName     string
	CustomerName  string
	CustomerEmail string
	CustomerPhone string
}

// CreateOrder creates a new order and returns orderID
func CreateOrder(_order Order) (*Order, error) {
	order := _order

	collection := client.Database("oms").Collection("orders")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, _order)
	if err != nil {
		log.Printf("CreateOrder:InsertOne %v\n", err)
		return nil, err
	}
	order.ID = result.InsertedID.(primitive.ObjectID)

	return &order, nil
}

// GetOrders returns list of order
func GetOrders(searchParam SearchParam) ([]Order, error) {
	var orders []Order

	collection := client.Database("oms").Collection("orders")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	// construct filter params
	filter := bson.M{}
	if searchParam.HotelName != "" {
		filter["hotelName"] = bson.M{"$regex": primitive.Regex{Pattern: "^.*" + searchParam.HotelName + ".*", Options: "i"}}
	}
	if searchParam.CustomerName != "" {
		filter["customer.name"] = bson.M{"$regex": primitive.Regex{Pattern: "^.*" + searchParam.CustomerName + ".*", Options: "i"}}
	}
	if searchParam.CustomerEmail != "" {
		filter["customer.email"] = bson.M{"$eq": searchParam.CustomerEmail}
	}
	if searchParam.CustomerPhone != "" {
		filter["customer.phone"] = bson.M{"$eq": searchParam.CustomerPhone}
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Printf("GetOrders:Find %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var order Order

		err := cursor.Decode(&order)
		if err != nil {
			log.Printf("GetOrders:Decode %v\n", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("GetOrders:Cursor %v\n", err)
		return nil, err
	}

	return orders, nil
}
