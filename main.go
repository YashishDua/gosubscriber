package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gosubscriber/database"
	"gosubscriber/endpoints"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	database.Init()
	//defer database.DB.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	// healthCheck
	router.HandleFunc("/health", GetHealth).Methods("GET")

	// endpoints
	router.HandleFunc("/getSubscribers", endpoints.GetSubscribers).Methods("GET")
	router.HandleFunc("/getSubscriber/{email}", endpoints.GetSubscriber).Methods("GET")
	router.HandleFunc("/addSubscriber", endpoints.AddSubscriber).Methods("POST")
	router.HandleFunc("/updateSubscriber/{email}", endpoints.UpdateSubscriber).Methods("POST")
	log.Fatal(http.ListenAndServe(":" + port, router))
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Get set Gooooooo!")
}
