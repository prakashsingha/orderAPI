package services

import "go.mongodb.org/mongo-driver/bson/primitive"

// Room structure
type Room struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	NoOfGuest int32              `json:"noOfGuest,omitempty" bson:"noOfGuest,omitempty"`
}
