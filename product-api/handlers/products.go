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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"my-go-project/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// swagger:parameters deleteProduct updateProduct
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

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Server is unable to return json", http.StatusInternalServerError)
	}
}

// swagger:route POST /products products addProduct
// Adds a new product
// responses:
//   201: productResponse
//   422: validationErrorResponse
//   501: errorResponse

// AddProduct adds a product to the data store
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)

}

// swagger:route PUT /products/{id} products updateProduct
// Updates an existing product
// responses:
//
//	204: noContent
//	422: validationErrorResponse
//	501: errorResponse
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	p.l.Println("Handle PUT Product", id)

	product := r.Context().Value(KeyProduct{}).(*data.Product)

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, product)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product
// responses:
//  204: noContent
//	501: errorResponse

// DeleteProduct deletes a product from the data store
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	p.l.Println("Handle DELETE Product", id)

	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
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

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := json.NewDecoder(r.Body).Decode(product)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
