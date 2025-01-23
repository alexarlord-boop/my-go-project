package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my-go-project/product-api/data"
	"net/http"
)

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
