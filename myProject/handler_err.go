package main

import (
	"net/http"
)

// readinessErrorHandler is a function that is used to handle the readiness of the application
// It accepts standard http.ResponseWriter and *http.Request parameters
// When called, it responds with a 400 Bad Request status code and a JSON payload containing {"error": "Something went wrong"}
// This endpoint can be used to test the error handling of the application
func readinessErrorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}

