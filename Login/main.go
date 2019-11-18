package main

import (
	"Login/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", server.SignupHandler).
		Methods("POST")
	/*r.HandleFunc("/login", server.LoginHandler).
		Methods("POST")
	r.HandleFunc("/users", server.UsersHandler).
		Methods("GET")*/

	log.Fatal(http.ListenAndServe(":8080", r))
}