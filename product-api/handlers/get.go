package handlers

import (
	"context"
	"net/http"

	protos "github.com/daffaromero/gomicro/currency/protos/currency"
	"github.com/daffaromero/gomicro/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")
	w.Header().Add("Content-Type", "application/json")

	prods := data.GetProducts()

	err := data.ToJSON(prods, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}

// swagger:route GET /products/{id} products listSingleProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET requests
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id := getProductID(r)

	p.l.Println("[DEBUG] get record id", id)

	prod, err := data.GetProductByID(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	} else if err != nil {
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	}

	if prod == nil {
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)
		return
	}

	// get exchange rate
	rr := &protos.RateRequest{
		Base:        string(protos.Currencies(protos.Currencies_value["CURRENCIES_EUR_UNSPECIFIED"])),
		Destination: string(protos.Currencies(protos.Currencies_value["CURRENCIES_GBP_UNSPECIFIED"])),
	}

	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("[Error] error getting new rate", err)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	p.l.Printf("Rate %#v", resp.Rate)

	prod.Price = prod.Price * resp.Rate

	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
