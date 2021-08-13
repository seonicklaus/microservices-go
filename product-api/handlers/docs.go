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

import "github.com/seonicklaus/microservices-go/product-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error messages retuned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products returns in the repsonse
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct{}

// swagger:parameters updateProduct createProduct
type productParameterWrapper struct {
	// Product data structure to Update or Create
	// Note: the id is ignored by the update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The ID of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}
