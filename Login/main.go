package main

import (
	"Login/server"
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
	r.HandleFunc("/signup", server.SignupHandler).
		Methods("POST")
	r.HandleFunc("/login", server.LoginHandler).
		Methods("POST")
	r.HandleFunc("/users", server.UsersHandler(formatter)).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}