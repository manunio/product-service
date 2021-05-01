// Package classification of Product API
//
// Documentation for product api
//
//    Scheme: http
//    BasePath: /
//    Version: 1.0.0
//
//    Consumes:
//    - application/json
//    Produces:
//    - application/json
// swagger:meta
package handlers

import "product-service/data"

// Note: Types defined here are purely for documentation purpose
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of errors
	// in: body
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All current products
	// in:body
	Body []data.Product
}

// Data structure representing single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponse struct {
}

// swagger:parameters updateResponse createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
