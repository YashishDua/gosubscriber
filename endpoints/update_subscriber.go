package endpoints

import "net/http"

func UpdateSubscriber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
