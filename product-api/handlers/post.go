package handlers

import (
	"net/http"

	"github.com/seonicklaus/microservices-go/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
// 	422: errorValidation
//	501: errorResponse

// Create handles POST request to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
