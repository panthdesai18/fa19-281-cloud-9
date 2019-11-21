package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"github.com/unrolled/render"
)

func main() {

	// port := os.Getenv("PORT")
	// if len(port) == 0 {
	// 	port = "12345"
	// }

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://cmpe281:cmpe281@3.89.47.220:27017")
	clientOptions := options.Client().ApplyURI(mongodb_server)

	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Mongo has been Connected")
	router := mux.NewRouter()

	router.HandleFunc("/placeOrder", CreateOrder).Methods("POST")
	router.HandleFunc("/removeOrder/{id}", RemoveOrder).Methods("DELETE")
	router.HandleFunc("/orders", GetOrders).Methods("GET")
	router.HandleFunc("/getUserOrder/{id}", GetOrderByUserId).Methods("GET")

	http.ListenAndServe(":3000", router)
}
