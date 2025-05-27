package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")
	godotenv.Load(".env")
	portString:= os.Getenv("PORT")
	if portString == ""{
		// log.Fatal is used to stop the program and print an error message
		log.Fatal("Port is not found in environment")
	}
	router := chi.NewRouter()
	// chi is a web framework that is used to create a router
	// chi.NewRouter() is used to create a new router

	// cors is a middleware that is used to allow cross-origin requests
	router.Use(cors.Handler(
		cors.Options{
			AllowedOrigins: []string{"https://*", "http://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"link"},
			AllowCredentials: false,
			MaxAge: 300,
		},
	))

	// srv is a server that is used to start the server
	srv := &http.Server{
		Handler: router,
		Addr: ":" + portString,
	}
	fmt.Printf("Starting server on port %v\n", portString)
	// srv.ListenAndServe() is used to start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	
	fmt.Println("Port: ", portString)
}
