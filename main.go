package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), http.FileServer(http.Dir("./"))))
}


