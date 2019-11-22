package controller

import (
	"context"
	"encoding/json"
	"Payments/config/db"
	"Payments/model"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateOrderPaymentHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var payment model.OrderPayment
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &payment)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetPaymentCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.OrderPayment 
	err = collection.FindOne(context.TODO(), bson.D{{"orderid", payment.OrderId}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {

			_, err = collection.InsertOne(context.TODO(), payment)
			if err != nil {
				res.Error = "Error While Creating Order Payment, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Order Payment Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Order payments already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

func GetOrderPayementHandler(w http.ResponseWriter, r *http.Request) () {

	var res model.ResponseResult
	collection, err := db.GetPaymentCollection()
	if err != nil {
		log.Fatal(err)
		fmt.Println("collection error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result model.OrderPayment
	orderPaymentId := mux.Vars(r)["orderId"]
	err = collection.FindOne(context.TODO(), bson.D{{"orderId", orderPaymentId}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Order Payment error")
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", result)

	json.NewEncoder(w).Encode(result)
	return
}

