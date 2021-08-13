package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l: l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello Requests")

	// read the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)

		http.Error(w, "Unable to read request body", http.StatusBadRequest)
	}

	// write the response
	fmt.Fprintf(w, "Hello, %s\n", body)
}
