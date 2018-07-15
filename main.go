package main

import (
	"log"
	"net/http"
	
	
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe("localhost:$PORT", http.FileServer(http.Dir("./"))))
}


