package handlers

import (
	"net/http"
	"product-service/data"
)

// swagger:route PUT /products products updateProduct
// Update a product details
//
// responses:
//		201: noContentResponse
// 404: errorResponse
// 422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {

	// fetch the product from the context
	product := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("[DEBUG] updating record id", product.ID)

	err := data.UpdateProduct(product)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		rw.WriteHeader(http.StatusNotFound)
		_ = data.ToJSON(&GenericError{Message: "Product not found in database"}, rw)
		return
	}

	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
