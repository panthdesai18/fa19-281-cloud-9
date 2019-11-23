package controller

import (
	"context"
	"encoding/json"
	"fa19-281-cloud-9/Locations/config/db"
	"fa19-281-cloud-9/Locations/model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
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

func GetALocationHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		fmt.Printf("Found a document: %+v\n", result)

		formatter.JSON(w, http.StatusOK, result)
		return
	}
}

func GetAllLocationHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res model.ResponseResult
		var results []*model.Location
		collection, err := db.GetDBLocationCollection()
		if err != nil {
			fmt.Println("collection error")
			formatter.JSON(w, http.StatusInternalServerError, "Internal Server Error")
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
		if results == nil || len(results)<=0 {
			formatter.JSON(w, http.StatusNotFound, "No restaurants found at the chosen location")
		} else {
			formatter.JSON(w, http.StatusOK, results)
		}
		fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
		//json.NewEncoder(w).Encode(results)
		return
	}
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
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

type ZipCodesResponse struct {
	Zip_codes []string `json:"zip_codes"`
}

func getZipCodes(body []byte) (*ZipCodesResponse, error) {
	var zip = new(ZipCodesResponse)
	err := json.Unmarshal(body, &zip)
	if(err != nil){
		fmt.Println("error:", err)
	}
	return zip, err
}

func GetLocationsByZipcodeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res model.ResponseResult
		var results []*model.Location
		setupResponse(&w, r)
		collection, err := db.GetDBLocationCollection()
		if err != nil {
			fmt.Println("collection error")
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		zipcode := mux.Vars(r)["zipcode"]
		resp, err := http.Get("https://www.zipcodeapi.com/rest/gZA5DC6sorrMU4qOSdSCXqjuB2ixwXPl6ERebFAVHMbf3Vy9KetjLrYnr6qo6qY6/radius.json/"+ zipcode +"/1/miles?minimal")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}


		log.Println(string(body))
		s, err := getZipCodes([]byte(body))

		fmt.Println(s.Zip_codes)
		if len(s.Zip_codes) > 0 {
			for index, each := range s.Zip_codes {
				fmt.Printf("Zip code [%d] is [%s]\n", index, each)
				cursor, err := collection.Find(context.TODO(), bson.D{{"zipcode", each}})
				if err != nil {
					log.Fatal(err)
				}

				for cursor.Next(context.TODO()) {
					var result model.Location
					err := cursor.Decode(&result)
					if err != nil {
						log.Fatal(err)
					}
					results = append(results, &result)
					fmt.Println(result)
				}
			}
		} else {
		cursor, err := collection.Find(context.TODO(), bson.D{{"zipcode", zipcode}})
			if err != nil {
				log.Fatal(err)
			}
			for cursor.Next(context.TODO()) {
				var result model.Location
				err := cursor.Decode(&result)
				if err != nil {
					log.Fatal(err)
				}
				results = append(results, &result)
				fmt.Println(result)
			}
		}

		fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
		formatter.JSON(w, http.StatusOK, results)
		return
	}
}
