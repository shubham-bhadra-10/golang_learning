package main

import (
	"net/http"
)

// readinessHandler is a function that is used to handle the readiness of the application
// It accepts standard http.ResponseWriter and *http.Request parameters
// When called, it responds with a 200 OK status code and a JSON payload containing {"status": "ok"}
// This endpoint can be used by load balancers and monitoring tools to verify the service is up and ready to accept requests

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}

