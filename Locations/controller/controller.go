package controller

import (
	"context"
	"encoding/json"
	"fa19-281-cloud-9/Locations/config/db"
	"fa19-281-cloud-9/Locations/model"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

func RegisterLocationHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var location model.Location
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &location)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBLocationCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.Location
	err = collection.FindOne(context.TODO(), bson.D{{"locationname", location.LocationName}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {

			_, err = collection.InsertOne(context.TODO(), location)
			if err != nil {
				res.Error = "Error While Creating Locations document, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Location Insertion Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Location document already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

func GetALocationHandler(w http.ResponseWriter, r *http.Request) () {

	var res model.ResponseResult
	collection, err := db.GetDBLocationCollection()
	if err != nil {
		//log.Fatal(err)
		fmt.Println("collection error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.Location
	locationId := mux.Vars(r)["locationId"]
	err = collection.FindOne(context.TODO(), bson.D{{"locationid", locationId}}).Decode(&result)
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Locations document error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)

	json.NewEncoder(w).Encode(result)
	return
}

func GetAllLocationHandler(w http.ResponseWriter, r *http.Request) () {

	var res model.ResponseResult
	var results []*model.Location
	collection, err := db.GetDBLocationCollection()
	if err != nil {
		fmt.Println("collection error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	for cursor.Next(context.TODO()) {
		var result model.Location
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
		fmt.Println(result)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	json.NewEncoder(w).Encode(results)
	return
}

func DeleteALocationHandler(w http.ResponseWriter, r *http.Request) () {

	var res model.ResponseResult
	collection, err := db.GetDBLocationCollection()
	if err != nil {
		//log.Fatal(err)
		fmt.Println("collection error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.Location
	locationId := mux.Vars(r)["locationId"]
	err = collection.FindOne(context.TODO(), bson.D{{"locationid", locationId}}).Decode(&result)
	if err == nil {
		collection.DeleteOne(context.TODO(), bson.D{{"locationid", locationId}})
		res.Result = "Document deleted successfully"
		json.NewEncoder(w).Encode(res)
		return
	}
	if err != nil {
		//log.Fatal(err)
		fmt.Println("Locations document error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(result)
	return
}

func GetLocationsByZipcodeHandler(w http.ResponseWriter, r *http.Request) () {

	var res model.ResponseResult
	var results []*model.Location
	collection, err := db.GetDBLocationCollection()
	if err != nil {
		fmt.Println("collection error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	zipcode := mux.Vars(r)["zipcode"]
	cursor, err := collection.Find(context.TODO(), bson.D{{"zipcode", zipcode}})
	for cursor.Next(context.TODO()) {
		var result model.Location
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &result)
		fmt.Println(result)
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	json.NewEncoder(w).Encode(results)
	return
}