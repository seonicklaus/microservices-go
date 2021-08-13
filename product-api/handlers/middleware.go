package handlers

import (
	"context"
	"net/http"

	"github.com/seonicklaus/microservices-go/product-api/data"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println(w, "[ERROR] deserializing product", http.StatusBadRequest)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		//validate the product
		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println(w, "[ERROR] validating product", errs)

			// return the validation message as an array
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(w, req)
	})
}
