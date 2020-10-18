package services

import "go.mongodb.org/mongo-driver/bson/primitive"

// Customer structure
type Customer struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Email string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone string             `json:"phone,omitempty" bson:"phone,omitempty"`
}
