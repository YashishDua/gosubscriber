package endpoints

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gosubscriber/model"
	"log"
	"net/http"
)

func GetSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	sub, fetchError := model.Get(params["email"])
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
		return
	}
	json.NewEncoder(w).Encode(sub)
}


func GetSubscribers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	subs, fetchError := model.GetAll()
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
		return
	}
	json.NewEncoder(w).Encode(subs)
}

