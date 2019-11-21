package main

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	ID          primitive.ObjectID `json:"_id,omitempty"    bson:"_id,omitempty"`
	ItemId      string             `json:"ItemId,omitempty" bson:"ItemId,omitempty"`
	ItemName    string             `json:"ItemName,omitempty" bson:"ItemName,omitempty"`
	Price       string             `json:"Price,omitempty" bson:"Price,omitempty"`
	Description string             `json:"Description,omitempty" bson:"Description,omitempty"`
	ItemType    string             `json:"ItemType,omitempty" bson:"ItemType,omitempty"`
}
