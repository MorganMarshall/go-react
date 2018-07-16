package main

import (
	
	"net/http"
	
)

func main() {


	// Simple static webserver:
	http.ListenAndServe(":$PORT", http.FileServer(http.Dir("public")))
}


