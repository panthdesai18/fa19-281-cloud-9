package main

import (
	"Payments/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/payments", controller.CreateOrderPaymentHandler).
		Methods("POST")
	r.HandleFunc("/payments/{orderId}", controller.GetOrderPayementHandler).
		Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
