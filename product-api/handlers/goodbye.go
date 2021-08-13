package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}

// newGoodbye creates a new Goodbye handler with a given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l: l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle Goodbye Requests")

	// write response
	fmt.Fprintf(w, "Goodbyeee\n")
}
