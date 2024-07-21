package main

import (
	"fmt"
	"testing"

	"github.com/daffaromero/gomicro/product-api/sdk/client"
	"github.com/daffaromero/gomicro/product-api/sdk/client/products"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:8086")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", prod.GetPayload()[0])
	t.Fail()
}
