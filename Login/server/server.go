package server

import (
	"Login/db"
	"Login/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/unrolled/render"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user types.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res types.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	collection, err := db.GetDBCollection()

	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}
	var result types.User
	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)

			if err != nil {
				res.Error = "Error While Hashing Password, Try Again"
				json.NewEncoder(w).Encode(res)
				return
			}
			user.Password = string(hash)

			_, err = collection.InsertOne(context.TODO(), user)
			if err != nil {
				res.Error = "Error While Creating User, Try Again!"
				json.NewEncoder(w).Encode(res)
				return
			}
			res.Result = "Signup Successful"
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "User already Exists!!"
	json.NewEncoder(w).Encode(res)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user types.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	collection, err := db.GetDBCollection()

	if err != nil {
		log.Fatal(err)
	}
	var result types.User
	var res types.ResponseResult

	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

	if err != nil {
		res.Error = "Invalid username"
		json.NewEncoder(w).Encode(res)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

	if err != nil {
		res.Error = "Invalid password"
		json.NewEncoder(w).Encode(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  result.Username,
		"fullname": result.FullName,
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		res.Error = "Error while generating token,Try again"
		json.NewEncoder(w).Encode(res)
		return
	}

	result.Token = tokenString
	//fmt.Println("Login successful!")
	result.Password = ""

	json.NewEncoder(w).Encode(result)
}

func UsersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res types.ResponseResult
		var results []*types.User
		collection, err := db.GetDBCollection()
		if err != nil {
			fmt.Println("Connection Error")
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}
		cursor, err := collection.Find(context.TODO(), bson.D{})
		for cursor.Next(context.TODO()) {
			var result types.User
			err := cursor.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &result)
			fmt.Println(result)
		}

		fmt.Printf("Found multiple documents: %+v\n", results)
		formatter.JSON(w, http.StatusOK, results)
		return
	}
}