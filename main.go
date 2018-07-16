package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
    if port == "" {
        port = "$PORT"
    }
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("./"))))
}


