package handlers

import (
	"encoding/json"
	"log"
	"my-go-project/product-api/data"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	encodedData, err := json.Marshal(productList)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
	}

	w.Write(encodedData)
}
