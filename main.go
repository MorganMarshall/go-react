package main

import (
	"log"
	"net/http"
	"os"
	
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(GetPort(), http.FileServer(http.Dir("./src"))))
}


func GetPort() string {
	 	var port = os.Getenv("PORT")
		// Set a default port if there is nothing in the environment
	 	if port == "" {
 		port = "4747"
	 	}
	 	return ":" + port
	 }