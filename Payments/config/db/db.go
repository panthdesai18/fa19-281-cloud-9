package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)


func GetPaymentCollection() (*mongo.Collection, error){

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@54.187.23.147:27017/admin?connect=direct"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("restaurant").Collection("payments")

	return collection, nil
}
