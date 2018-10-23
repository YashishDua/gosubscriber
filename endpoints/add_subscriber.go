package endpoints

import (
	"encoding/json"
	"gosubscriber/model"
	"log"
	"net/http"
)

func AddSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sub model.Subscriber
	json.NewDecoder(r.Body).Decode(&sub)
	fetchError := sub.Add()
	if fetchError != nil {
		log.Printf("Error while fetching from DB: %s", fetchError)
		return
	}
	json.NewEncoder(w).Encode(sub)
}

