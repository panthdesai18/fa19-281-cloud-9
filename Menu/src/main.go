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

	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://cmpe281:cmpe281@3.89.47.220:27017")
	clientOptions := options.Client().ApplyURI(mongodb_server)
	fmt.Println("Client Options set...")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Mongo Connected...")
	router := mux.NewRouter()

	router.HandleFunc("/item", CreateMenuEndpoint).Methods("POST")
	router.HandleFunc("/item/{id}", RemoveItemEndpoint).Methods("DELETE")
	router.HandleFunc("/items", GetMenuItemsEndpoint).Methods("GET")
	router.HandleFunc("/item/{id}", GetMenuEndpoint).Methods("GET")

	http.ListenAndServe(":12345", router)
}
