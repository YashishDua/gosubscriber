package api

import (
	db "github.com/gosubscriber/database"
	"github.com/gosubscriber/model"
	"encoding/json"
	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var databaseReference *pg.DB

func Init() {
	// Connect to database
	databaseReference = db.Connect()
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Get set Gooooooo!")
}

func AddSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sub model.Subscriber
	json.NewDecoder(r.Body).Decode(&sub)
	fetchError := sub.Add(databaseReference)
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
	}
	json.NewEncoder(w).Encode(sub)
}

func GetSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sub, fetchError := model.Get(databaseReference, params["email"])
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
		return
	}
	json.NewEncoder(w).Encode(sub)
}

func GetSubscribers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	subs, fetchError := model.GetAll(databaseReference)
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
		return
	}
	json.NewEncoder(w).Encode(subs)
}

func UpdateSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
