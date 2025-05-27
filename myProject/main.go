package main

import (
	"fmt"
	"log"
	"os"

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
	fmt.Println("Port: ", portString)
}
