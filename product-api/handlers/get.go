package handlers

import (
	"net/http"

	"github.com/seonicklaus/microservices-go/product-api/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	w.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, w)
	if err != nil {
		//we should never get here but log the error just in case
		p.l.Println("[ERROR] serializing product, ", err)
	}
}

// swagger:route GET /product/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
// 	404: errorResponse

// ListSingle handles GET request
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id ", id)
	prod, err := data.GetProductByID(id)
	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product ", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	default:
		p.l.Println("[ERROR] fetching product ", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		// we should never get here but log the error just in case
		p.l.Println("[ERROR] serializing product, ", err)
	}
}
