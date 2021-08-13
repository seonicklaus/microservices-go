package handlers

import (
	"net/http"

	"github.com/seonicklaus/microservices-go/product-api/data"
)

// swagger:route PUT /products products updateProduct
// Update a product details
//
// responses:
//	201: noContentResponse
// 	404: errorResponse
//	422: errorValidation

// Update handles PUT requests to update products
func (p Products) Update(w http.ResponseWriter, r *http.Request) {
	//fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record id ", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	// Write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
