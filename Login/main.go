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
	r.HandleFunc("/adminlogin", server.LoginHandler).
		Methods("POST")
	r.HandleFunc("/users", server.GetAllUsers(formatter)).
		Methods("GET")
	r.HandleFunc("/user/{username}", server.GetOneUser(formatter)).
		Methods("GET")
	r.HandleFunc("/user/{emailid}", server.GetOneUserByEmailId(formatter)).
		Methods("GET")
	r.HandleFunc("/user/{username}", server.DeleteAUser).
		Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}