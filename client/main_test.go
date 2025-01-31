package main

import (
	"my-go-project/client/client"
	"my-go-project/client/client/products"
	"testing"

	"github.com/go-openapi/strfmt"
)

func TestOurClient(t *testing.T) {
	if client.Default == nil {
		t.Error("client.Default is nil")
	} else {
		t.Log("Our client is working")
	}
}

// had to fix response header in original API - add content-type application/json
func TestListProducts(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	products, err := c.Products.ListProducts(products.NewListProductsParams())
	if err != nil {
		t.Fatal(err)
	}

	if len(products.Payload) == 0 {
		t.Error("Expected products, got none")
	}
}

func TestGetProduct(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8080")
	c := client.NewHTTPClientWithConfig(strfmt.Default, cfg)

	product, err := c.Products.ListSingleProduct(products.NewListSingleProductParams().WithID(1))
	if err != nil {
		t.Fatal(err)
	}

	if product.Payload == nil {
		t.Error("Expected product, got none")
	}
}
