package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	PaymentPending int = iota
	PaymentComplete
)

var (
	Status = map[string]string{
		"0": "Pending",
		"1": "Complete",
	}
)

// Payment is for payment details
type Payment struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Status        int                `json:"status" bson:"status"`
	ConfirmedDate *time.Time         `json:"confirmedDate,omitempty" bson:"confirmedDate,omitempty"`
	CardNo        string             `json:"cardNo,omitempty" bson:"cardNo,omitempty"`
}

// CreatePayment creates a payment with PENDING status
func CreatePayment(_payment Payment) (string, error) {
	collection := client.Database("oms").Collection("payments")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, _payment)
	if err != nil {
		log.Printf("CreatePayment:InsertOne %v\n", err)
		return "", err
	}

	oID := result.InsertedID.(primitive.ObjectID)
	if oID.IsZero() {
		return "", err
	}
	return oID.Hex(), nil
}

// GetPayment returns single result of payment
func GetPayment(_id string) (*Payment, error) {
	var payment Payment

	paymentID, paymentErr := primitive.ObjectIDFromHex(_id)
	if paymentErr != nil {
		return nil, fmt.Errorf("GetPayment:ObjectIDFromHex %v\n", paymentErr.Error())
	}

	collection := client.Database("oms").Collection("payments")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"_id": paymentID})
	err := result.Decode(&payment)
	if err != nil {
		return nil, fmt.Errorf("GetPaymentStatus:Decode %v \n", err.Error())
	}
	return &payment, nil
}

// GetPaymentDetail returns payment status
func GetPaymentDetail(_id string) (string, error) {
	status := ""
	payment, err := GetPayment(_id)
	if err != nil {
		log.Printf("GetPaymentStatus:Decode %v \n", err)
		return status, err
	}

	if payment == nil {
		return "", nil
	}

	statusStr := strconv.Itoa(payment.Status)
	return Status[statusStr], nil
}

func UpdatePayment(_payment Payment) (int64, error) {
	collection := client.Database("oms").Collection("payments")
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": _payment.ID},
		bson.D{
			{"$set", bson.D{{"status", _payment.Status}}},
			{"$set", bson.D{{"confirmedDate", time.Now()}}},
			{"$set", bson.D{{"cardNo", _payment.CardNo}}},
		})

	if err != nil {
		log.Printf("UpdatePayment:UpdateOne %v\n", err)
		return 0, err
	}

	return result.ModifiedCount, nil
}
