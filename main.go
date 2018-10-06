package main

import (
	endPoints "./rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	endPoints.Init()

	// healthCheck
	router.HandleFunc("/api/health", endPoints.GetHealth).Methods("GET")

	// endpoints
	router.HandleFunc("/api/getSubscribers", endPoints.GetSubscribers).Methods("GET")
	router.HandleFunc("/api/getSubscriber/{email}", endPoints.GetSubscriber).Methods("GET")
	router.HandleFunc("/api/addSubscriber", endPoints.AddSubscriber).Methods("POST")
	router.HandleFunc("/api/updateSubscriber/{email}", endPoints.UpdateSubscriber).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))

}

