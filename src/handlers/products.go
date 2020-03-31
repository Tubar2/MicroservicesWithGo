package handlers

import (
	"BuildingMicroservicesWithGo_NicJackson/src/data"
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

// PutProduct implementats the PUT operation on products
// It updates the databasse
func (p *Products) PutProduct(respW http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(respW, "Unable to parse URL", http.StatusBadRequest)
		return
	}

	p.l.Println("Handling PUT operation")

	// create empty product
	prod := &data.Product{}

	// decodes request json into empty product
	err = prod.FromJSON(req.Body)
	if err != nil {
		http.Error(respW, "Unable to decode json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(respW, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(respW, "Product not found", http.StatusInternalServerError)
		return
	}

}

// PostProduct implementats the POST operation on product
// It adds a product to the databse
func (p *Products) PostProduct(respW http.ResponseWriter, req *http.Request) {
	p.l.Println("Handling POST operation")

	// create empty product
	prod := &data.Product{}

	// decodes request json into empty product
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(respW, "Unable to decode json", http.StatusBadRequest)
		return
	}

	// append new product into database
	data.AddProduct(prod)

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
		return
	}
}
