package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
	  return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
  }

func main() {
	// Simple static webserver:
	addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }


	log.Fatal(http.ListenAndServe(addr, http.FileServer(http.Dir("./"))))
}


