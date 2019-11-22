package main

import (
	"fa19-281-cloud-9/Locations/controller"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	r.HandleFunc("/locations", controller.RegisterLocationHandler).
		Methods("POST")
	r.HandleFunc("/locations/{locationId}", controller.GetALocationHandler(formatter)).
		Methods("GET")
	r.HandleFunc("/locations", controller.GetAllLocationHandler(formatter)).
		Methods("GET")
	/*r.HandleFunc("/locations/{locationId}", controller.UpdateALocationHandler).
		Methods("PUT")*/
	r.HandleFunc("/locations/{locationId}", controller.DeleteALocationHandler).
		Methods("DELETE")
	r.HandleFunc("/locations/zipcode/{zipcode}", controller.GetLocationsByZipcodeHandler(formatter)).
		Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}
