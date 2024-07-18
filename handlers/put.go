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
	w.Header().Add("Content-Type", "application/json")
	id := getProductID(r)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Println("[DEBUG] Updating product with id", id)

	err := data.UpdateProduct(prod, id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return

	} else if err != nil {
		p.l.Println("[ERROR] updating record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	data.ToJSON(prod, w)
}
