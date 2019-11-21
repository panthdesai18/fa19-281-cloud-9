package main

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"time"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gorilla/mux"
)

var client *mongo.Client

var mongodb_server = "mongodb://root:password@10.0.1.89:27017"

// var mongodb_server = "mongodb://localhost:27017"
var mongodb_server_1 = "52.37.128.85:27017"
var mongodb_database = "Order"
var mongodb_collection = "ORDER"
var mongodb_collection1 = "submissions"
var mongodb_username = "root"
var mongodb_password = "password"

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CreateOrder(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside Create function")

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	var order Order
	_ = json.NewDecoder(request.Body).Decode(&order)
	fmt.Println("Order Amount", order.TotalAmount)
	collection := client.Database(mongodb_database).Collection("ORDER")
	fmt.Println("Order Amount", order.TotalAmount)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, order)
	json.NewEncoder(response).Encode(result)
}

func GetOrderByUserId(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["UserId"])
	var order Order
	collection := client.Database(mongodb_database).Collection("ORDER")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Order{OrderId: id}).Decode(&order)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(order)
}

func RemoveOrder(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var order Order

	collection := client.Database(mongodb_database).Collection("ORDER")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := Order{OrderId: id}
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println(result.DeletedCount)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(order)
}

func GetOrders(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------
	fmt.Println("Inside Get Orders function")
	response.Header().Set("content-type", "application/json")
	var orders []Order
	collection := client.Database(mongodb_database).Collection("ORDER")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var order Order
		cursor.Decode(&order)
		orders = append(orders, order)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(orders)
}
