package handlers

import (
	"my-go-project/product-api/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetList(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	productList := data.GetProducts()
	err := productList.ToJSON(w)
	if err != nil {
		http.Error(w, "Server is unable to return json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Returns a single product by ID
// responses:
//  200: productResponse

// GetProduct handles GET requests and returns a single product by ID
func (p *Products) GetDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle GET Product", id)

	product, err := data.GetProductByID(id)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	err = product.ToJSON(w)
	if err != nil {
		http.Error(w, "Server is unable to return json", http.StatusInternalServerError)
	}
}
