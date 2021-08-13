package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/seonicklaus/microservices-go/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the product list from the datastore
	lp := data.GetProducts()

	err := lp.ToJSON(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)
}

func (p Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println(w, "[ERROR] deserializing product", http.StatusBadRequest)
			http.Error(w, "Error reading product", http.StatusBadRequest)
			return
		}

		//validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println(w, "[ERROR] validating product", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product, %s\n", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
