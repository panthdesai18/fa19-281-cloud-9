package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var client *mongo.Client
var mongodb_server = "mongodb+srv://cmpe281:@cmpe281-ievds.mongodb.net/test?retryWrites=true&w=majority"
var mongodb_server_1 = "52.37.128.85:27017"

var mongodb_database = "Menu"
var mongodb_collection = "menu"

type Order struct {
	OrderId     string   `json:"OrderId,omitempty" bson:"OrderId,omitempty"`
	UserId      string   `json:"UserId,omitempty" bson:"UserId,omitempty"`
	OrderStatus string   `json:"OrderStatus,omitempty" bson:"OrderStatus,omitempty"`
	Items       []string `json:"Items,omitempty" bson:"Items,omitempty"`
	TotalAmount string   `json:"TotalAmount,omitempty" bson:"TotalAmount,omitempty"`
}

func CreateOrderEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"menu-shard-00-00-mgoh4.mongodb.net:27017",
			"menu-shard-00-01-mgoh4.mongodb.net:27017",
			"menu-shard-00-02-mgoh4.mongodb.net:27017"},
		Database: "admin",
		Username: "m_udit",
		Password: "cmpe281",
	}
	fmt.Println("After Connection String")
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	fmt.Println("After Connection")

	if err != nil {
		fmt.Println("Errorrrrr")
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var order Order
	_ = json.NewDecoder(request.Body).Decode(&order)
	err = c.Insert(order)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode("OK")
}

func GetOrderByUserId(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"menu-shard-00-02-mgoh4.mongodb.net:27017",
			"menu-shard-00-01-mgoh4.mongodb.net:27017",
			"menu-shard-00-00-mgoh4.mongodb.net:27017"},
		Database: "admin",
		Username: "m_udit",
		Password: "cmpe281",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var order []Order

	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	fmt.Println(id)
	err = c.Find(bson.M{"UserId": id}).All(&order)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode(order)
}

type smtpServer struct {
	host string
	port string
}

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//clientOptions := options.Client().ApplyURI("mongodb://cmpe281:cmpe281@3.89.47.220:27017")
	clientOptions := options.Client().ApplyURI("mongodb+srv://jay:jay@movies-upn2q.mongodb.net/test?retryWrites=true&w=majority")
	fmt.Println("Client Options set...")
	client, _ = mongo.Connect(ctx, clientOptions)
	fmt.Println("Mongo has been Connected")
	router := mux.NewRouter()
	router.HandleFunc("/order", CreateOrderEndpoint).Methods("POST")
	router.HandleFunc("/order/{id}", GetOrderByUserId).Methods("GET")
	http.ListenAndServe(":3000", router)
}
