package handlers

import (
	"net/http"
	"product-service/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//		200: productResponse
// 422: errorValidation
// 501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	product := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %v#\n", product)
	data.AddProduct(product)
}
