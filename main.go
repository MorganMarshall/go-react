package main

import (
	"log"
	"net/http"
	
)

func main() {
	// Simple static webserver:
	port := os.Getenv("PORT")


	log.Fatal(http.ListenAndServe(port, http.FileServer(http.Dir("./"))))
}


