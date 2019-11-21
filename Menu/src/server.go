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

var mongodb_server = "mongodb://root:password@10.0.1.191:27017"

//var mongodb_server = "mongodb://localhost:27017"
var mongodb_server_1 = "52.37.128.85:27017"
var mongodb_database = "menu"
var mongodb_collection = "menu"
var mongodb_collection1 = "submissions"
var mongodb_username = "admin"
var mongodb_password = "admin"

// type Person struct {
// 	UserId string `json:"UserId,omitempty" bson:"UserId,omitempty"`
// 	Name   string `json:"name,omitempty" bson:"name,omitempty"`
// 	Email  string `json:"email,omitempty" bson:"email,omitempty"`
// 	Mobile string `json:"mobile,omitempty" bson:"mobile,omitempty"`
// }

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func CreateMenuEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside Create function")

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	var menu Menu
	_ = json.NewDecoder(request.Body).Decode(&menu)
	fmt.Println("ItemName", menu.ItemName)
	collection := client.Database(mongodb_database).Collection("menu")
	fmt.Println("ItemName", menu.ItemName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, menu)
	json.NewEncoder(response).Encode(result)
}

func GetMenuEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var menu Menu
	collection := client.Database(mongodb_database).Collection("menu")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Menu{ID: id}).Decode(&menu)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(menu)
}

func RemoveItemEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var menu Menu

	collection := client.Database(mongodb_database).Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	filter := Menu{ID: id}
	result, err := collection.DeleteOne(ctx, filter)
	fmt.Println(result.DeletedCount)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(menu)
}

func GetMenuItemsEndpoint(response http.ResponseWriter, request *http.Request) {

	//------------------LOCALHOST CODE--------------------

	response.Header().Set("content-type", "application/json")
	var menus []Menu
	collection := client.Database(mongodb_database).Collection("menu")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var menu Menu
		cursor.Decode(&menu)
		menus = append(menus, menu)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(menus)
}
