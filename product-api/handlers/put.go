package handlers

import (
	"my-go-project/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /products/{id} products updateProduct
// Updates an existing product
// responses:
//
//	204: noContent
//	422: validationErrorResponse
//	501: errorResponse
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
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
