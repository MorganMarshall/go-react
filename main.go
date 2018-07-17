package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Simple static webserver:
	port := os.Getenv("$PORT")

	log.Fatal(http.ListenAndServe(":8080" + port, http.FileServer(http.Dir("./"))))
}


