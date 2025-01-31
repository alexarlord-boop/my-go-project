// Product API
//
// # Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package handlers

import (
	"log"
	"my-go-project/product-api/data"
)

// A list of products
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in:body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponse struct {
	// The created product
	// in: body
	Body data.Product
}

// Validation errors as an array of strings
// swagger:response validationErrorResponse
type validationErrorResponse struct {
	// The error message
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// Generic error message as a string
// swagger:response errorResponse
type errorResponse struct {
	// The error message
	// in: body
	Body struct {
		Message string `json:"message"`
	}
}

// No content is returned by this API endpoint
// swagger:response noContent
type noContent struct {
	// No content
}

// swagger:parameters deleteProduct updateProduct listSingleProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the data store
	// in:path
	// required: true
	ID int `json:"id"`
}

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}
