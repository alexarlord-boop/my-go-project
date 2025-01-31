package handlers

import (
	"my-go-project/product-api/data"
	"net/http"
)

// swagger:route POST /products products addProduct
// Adds a new product
// responses:
//   201: productResponse
//   422: validationErrorResponse
//   501: errorResponse

// AddProduct adds a product to the data store
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)

}
