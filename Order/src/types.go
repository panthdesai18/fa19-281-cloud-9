package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Person struct {
	ID     primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name,omitempty"`
	Email  string             `json:"email,omitempty" bson:"email,omitempty"`
	Mobile string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
}

type Order struct {
	OrderId     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId      string             `json:"UserId,omitempty" bson:"UserId,omitempty"`
	OrderStatus string             `json:"OrderStatus,omitempty" bson:"OrderStatus,omitempty"`
	Items       []string           `json:"Items,omitempty" bson:"Items,omitempty"`
	TotalAmount string             `json:"TotalAmount,omitempty" bson:"TotalAmount,omitempty"`
}
