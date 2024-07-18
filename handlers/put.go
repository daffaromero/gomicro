package handlers

import (
	"gomicro/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] Updating product with id", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
