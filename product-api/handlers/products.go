package handlers

import (
	"context"
	"encoding/json"
	"log"
	"my-go-project/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Server is unable to return json", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	product := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(product)

}

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

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		product := &data.Product{}
		err := json.NewDecoder(r.Body).Decode(product)
		if err != nil {
			http.Error(w, "Unable to decode json", http.StatusBadRequest)
			return
		}

		// err = product.Validate()
		// if err != nil {
		// 	http.Error(w, "Product validation failed", http.StatusBadRequest)
		// 	return
		// }

		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
