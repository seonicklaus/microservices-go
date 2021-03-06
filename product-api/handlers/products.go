// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/seonicklaus/microservices-go/product-api/data"
)

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// Product handler for getting and updating products
type Products struct {
	l *log.Logger
	v *data.Validation
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("invalid Path, path should be /products/{id}")

// NewProducts returns a new product handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l: l, v: v}
}

// GenericError is a generic error message return by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the URL
	vars := mux.Vars(r)

	// convert the id into an integer amd return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
