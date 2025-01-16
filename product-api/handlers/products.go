package handlers

import (
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

// modification with ToJSON method -- simplies the code, making it a bit faster.
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Server is unable to return json", http.StatusInternalServerError)
	}

}
