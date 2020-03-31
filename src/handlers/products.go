package handlers

import (
	"BuildingMicroservicesWithGo_NicJackson/src/data"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products defines a handler for product
type Products struct {
	l *log.Logger
}

// NewProducts defines a function that returns a new Products object handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts implementats the GET operation on products
// It retrieves the database
func (p *Products) GetProducts(respW http.ResponseWriter, req *http.Request) {
	p.l.Println("Handling GET operation")

	// fetch products from databse
	lp := data.GetProducts()

	// parse fetched products to JSON
	err := lp.ToJSON(respW)
	if err != nil {
		http.Error(respW, "Unable to encode json", http.StatusInternalServerError)
	}
}

// PostProduct implementats the POST operation on product
// It adds a product to the databse
func (p *Products) PostProduct(respW http.ResponseWriter, req *http.Request) {
	p.l.Println("Handling POST operation")

	// create empty product
	prod := req.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)

}

// PutProduct implementats the PUT operation on products
// It updates the databasse
func (p *Products) PutProduct(respW http.ResponseWriter, req *http.Request) {
	// parse id from URL: localhost:9090/id
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(respW, "Unable to parse URL", http.StatusBadRequest)
		return
	}

	p.l.Println("Handling PUT operation")

	// create empty product
	prod := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(respW, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(respW, "Product not found", http.StatusInternalServerError)
		return
	}

}

// KeyProduct is
type KeyProduct struct{}

// MiddlewareValidateProduct updates request so it contains valid product object
func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(respW http.ResponseWriter, req *http.Request) {
		// create empty product
		prod := data.Product{}

		// decodes request json into product
		err := prod.FromJSON(req.Body)
		if err != nil {
			http.Error(respW, "Unable to decode json", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(respW, req)

	})
}
