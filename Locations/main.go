package main

import (
	"fa19-281-cloud-9/Locations/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Sample struct {
	A int `json:"a"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/locations", controller.RegisterLocationHandler).
		Methods("POST")
	r.HandleFunc("/locations/{locationId}", controller.GetALocationHandler).
		Methods("GET")
	/*r.HandleFunc("/locations", controller.GetAllLocationHandler).
		Methods("GET")
	r.HandleFunc("/locations/{locationId}", controller.UpdateALocationHandler).
		Methods("PUT")
	r.HandleFunc("/locations/{locationId}", controller.DeleteALocationHandler).
		Methods("DELETE")*/
	log.Fatal(http.ListenAndServe(":8080", r))
}


/*package cmpe281project

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		//clientOptions := options.Client().ApplyURI("mongodb://54.71.153.203:27017")
		fmt.Println("before : Connected to MongoDB!")
		// Connect to MongoDB
		//client, err := mongo.Connect(context.TODO(), clientOptions)
		// Initialize a new mongo client with options
		client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://admin:admin@54.71.153.203:27017/admin?connect=direct"))

		// Connect the mongo client to the MongoDB server
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err = client.Connect(ctx)

		if err != nil {
			log.Fatal(err)
			fmt.Println("error")
		}

		// Check the connection
		err = client.Ping(context.TODO(), nil)

		if err != nil {
			log.Fatal(err)
			fmt.Println("error")
		}

		fmt.Println("Connected to MongoDB!")
		collection := client.Database("test").Collection("test")
		cursor, err := collection.Find(context.TODO(), bson.D{})
		fmt.Println(cursor)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))

	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func cmpe281project() {
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

 */