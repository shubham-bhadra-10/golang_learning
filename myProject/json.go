package main

import (
	"encoding/json"
	"log"
	"net/http"
)
// respondWithJSON is a function that is used to respond with a JSON payload
func respondWithJSON(w http.ResponseWriter, code int, payload interface{})  {
	// Marshal is a function that is used to marshal a payload into a JSON payload
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		w.WriteHeader(500)
		return 
	}
	// Add is a function that is used to add a header to the response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)

}
