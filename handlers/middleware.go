package handlers

import (
	"context"
	"gomicro/data"
	"net/http"
)

// MiddlewareValidateProduct validates the product in the request and calls next if ok
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		errs := p.v.Validate(prod)
		if len(errs) != 0 {
			p.l.Println("[ERROR] validating product", errs)

			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
