package main

import (
	"log"
	"net/http"
	"os"
	
	
)

func main() {
	// Simple static webserver:
	
	log.Fatal(http.ListenAndServe("/", http.FileServer(http.Dir("./"))))
}


