package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"

	"net/http"
	"time"

	//"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"

	//"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

var client *mongo.Client
var mongodb_server = "mongodb+srv://cmpe281:@cmpe281-ievds.mongodb.net/test?retryWrites=true&w=majority"
var mongodb_server_1 = "52.37.128.85:27017"
var mongodb_database = "Menu"
var mongodb_collection = "menu"
var mongodb_collection1 = "submissions"
var mongodb_username = "panthdesai18"
var mongodb_password = "cmpe272"

type Menu struct {
	ItemId      string `json:"ItemId,omitempty" bson:"ItemId,omitempty"`
	ItemName    string `json:"ItemName,omitempty" bson:"ItemName,omitempty"`
	Price       string `json:"Price,omitempty" bson:"Price,omitempty"`
	Description string `json:"Description,omitempty" bson:"Description,omitempty"`
	ItemType    string `json:"ItemType,omitempty" bson:"ItemType,omitempty"`
}

func CreateMenuEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cmpe281-shard-00-02-ievds.mongodb.net:27017",
			"cmpe281-shard-00-01-ievds.mongodb.net:27017",
			"cmpe281-shard-00-00-ievds.mongodb.net:27017"},
		Database: "admin",
		Username: "cmpe281",
		Password: "cmpe281",
	}
	fmt.Println("Inside get assignemts function")
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	fmt.Println("Inside get assignemts function")
	//session, err := mgo.ParseURLDial(mongodb_server)
	if err != nil {
		fmt.Println("Inside get assignemts function")
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var menu Menu
	_ = json.NewDecoder(request.Body).Decode(&menu)
	err = c.Insert(menu)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode("OK")
	// var assignments_array []assignment
	// err = c.Find(bson.M{}).All(&assignments_array)
	// fmt.Println("Assignments", assignments_array)
	//formatter.JSON(w, http.StatusOK, assignments_array)

	// response.Header().Set("content-type", "application/json")
	// var person Person
	// _ = json.NewDecoder(request.Body).Decode(&person)
	// fmt.Println("Name", person.Name)
	// collection := client.Database("movies").Collection("people")
	// fmt.Println("Name", person.Name)
	// ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// result, _ := collection.InsertOne(ctx, person)
	// json.NewEncoder(response).Encode(result)
}

func GetMenuEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cmpe281-shard-00-00-ievds.mongodb.net:27017",
			"cmpe281-shard-00-01-ievds.mongodb.net:27017",
			"cmpe281-shard-00-02-ievds.mongodb.net:27017"},
		Database: "admin",
		Username: "panthdesai18",
		Password: "cmpe272",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	//session, err := mgo.ParseURLDial(mongodb_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var menu []Menu
	//_ = json.NewDecoder(request.Body).Decode(&person)
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := (params["id"])
	fmt.Println(id)
	err = c.Find(bson.M{"ItemId": id}).All(&menu)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode(menu)
}

func RemoveMenuEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cmpe281-shard-00-00-ievds.mongodb.net:27017",
			"cmpe281-shard-00-01-ievds.mongodb.net:27017",
			"cmpe281-shard-00-02-ievds.mongodb.net:27017"},
		Database: "admin",
		Username: "cmpe281",
		Password: "cmpe281",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	//session, err := mgo.ParseURLDial(mongodb_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	//var person []Person
	//_ = json.NewDecoder(request.Body).Decode(&person)
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := (params["id"])
	fmt.Println(id)
	err = c.Remove(bson.M{"ItemId": id})
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode("Deleted Record")
}

func GetMenusEndpoint(response http.ResponseWriter, request *http.Request) {

	fmt.Println("Inside get assignemts function")

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs: []string{"cmpe281-shard-00-00-ievds.mongodb.net:27017",
			"cmpe281-shard-00-01-ievds.mongodb.net:27017",
			"cmpe281-shard-00-02-ievds.mongodb.net:27017"},
		Database: "admin",
		Username: "cmpe281",
		Password: "cmpe281",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)

	//session, err := mgo.ParseURLDial(mongodb_server)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(mongodb_database).C(mongodb_collection)
	var menu []Menu
	//_ = json.NewDecoder(request.Body).Decode(&person)
	response.Header().Set("content-type", "application/json")
	//params := mux.Vars(request)
	// id, _ := (params["id"])
	// fmt.Println(id)
	err = c.Find(bson.M{}).All(&menu)
	if err != nil {
		fmt.Println("Can not insert")
	}
	json.NewEncoder(response).Encode(menu)
}

// smtpServer data to smtp server
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
	fmt.Println("Mongo Connected...")
	router := mux.NewRouter()
	router.HandleFunc("/menu", CreateMenuEndpoint).Methods("POST")
	router.HandleFunc("/menu/{id}", RemoveMenuEndpoint).Methods("DELETE")
	router.HandleFunc("/menus", GetMenusEndpoint).Methods("GET")
	router.HandleFunc("/menu/{id}", GetMenuEndpoint).Methods("GET")

	http.ListenAndServe(":12345", router)
}
