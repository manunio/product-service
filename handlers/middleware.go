package handlers

import (
	"context"
	"net/http"
	"product-service/data"
)

//  MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := &data.Product{}

		err := data.FromJSON(product, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			rw.WriteHeader(http.StatusBadRequest)
			_ = data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		// validate product
		errs := p.v.Validate(product)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", err)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			_ = data.ToJSON(&ValidationError{Messages: errs.Errors()}, rw)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// call the next handler , which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}
