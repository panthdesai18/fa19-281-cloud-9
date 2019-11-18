package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetDBLocationCollection() (*mongo.Collection, error){

	//clientOptions := options.Client().ApplyURI("mongodb://admin:admin@54.71.153.203:27017/admin?connect=direct")
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@54.71.153.203:27017/admin?connect=direct"))
	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		fmt.Println("error")
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	collection := client.Database("restaurant").Collection("locations")

	return collection, nil
}
