package handlers

import (
	"net/http"

	"github.com/daffaromero/gomicro/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(*data.Product)

	p.l.Printf("[DEBUG] Inserting product: %v\n", &prod)
	data.AddProduct(prod)

	err := data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
