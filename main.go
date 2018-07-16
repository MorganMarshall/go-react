package main

import (
	"log"
	"net/http"
	"os"
)

func main() {


	// Simple static webserver:
	log.Fatal(http.ListenAndServe("$PORT", http.FileServer(http.Dir("./"))))
}


