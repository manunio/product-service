package handlers

import (
	"net/http"
	"product-service/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products from the database
// responses:
//		200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	products := data.GetProducts()

	err := data.ToJSON(products, rw)
	if err != nil {
		// we should never be here but log error just in-case
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route /GET /products/{id} products listSingle
// Returns a list of products from the database
// responses:
//		200: productResponse
//		404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	product, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusNotFound)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		p.l.Println("[ERROR] fetching product", err)

		rw.WriteHeader(http.StatusInternalServerError)
		_ = data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = data.ToJSON(product, rw)
	if err != nil {
		// we should never be here but log the error just in-case
		p.l.Println("[ERROR] serializing product", err)
	}
}
