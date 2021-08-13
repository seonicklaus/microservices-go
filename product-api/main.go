package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// requests to the path /goodbye will be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other requests will be handled by this function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")

		// read the body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(w, "Error reading body", http.StatusBadRequest)
		}

		// write the response
		fmt.Fprintf(w, "Hello, %s\n", body)
	})

	// listen for connection on all ip addresses (0.0.0.0)
	// port: 9090
	log.Println("Starting server")
	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
