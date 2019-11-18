package main

import (
	"fa19-281-cloud-9/Locations/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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
