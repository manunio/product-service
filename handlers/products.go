package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"product-service/data"
	"strconv"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a product handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// getProducts returns the products from the data-store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the data-store
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// addProduct adds a product to the data-store
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)

}

// updateProducts updates a product to the data-store
func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT products", id)

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}
		err := product.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// call the next handler , which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}
