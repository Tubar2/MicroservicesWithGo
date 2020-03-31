package handlers

import (
	"BuildingMicroservicesWithGo_NicJackson/src/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Products defines a handler for product
type Products struct {
	l *log.Logger
}

// NewProducts defines a function that returns a new Products object handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler
func (p *Products) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	// handle Get operation requests
	if req.Method == http.MethodGet {
		p.getProducts(respW, req)
		return
	}

	// handle Post operation requests
	if req.Method == http.MethodPost {
		p.postProduct(respW, req)
		return
	}

	// handle Put operation requests
	if req.Method == http.MethodPut {
		p.putProduct(respW, req)
		return
	}

	//Catch all non implemented operations and send error
	respW.WriteHeader(http.StatusMethodNotAllowed)
}

// Implementation of the PUT operation on products
// Updates databasse
func (p *Products) putProduct(respW http.ResponseWriter, req *http.Request) {
	p.l.Println("Handling PUT operation")
	r := regexp.MustCompile(`/([0-9]+)`)
	g := r.FindAllStringSubmatch(req.URL.Path, -1)

	if len(g) != 1 {
		http.Error(respW, "Unable to decode json", http.StatusBadRequest)
		return
	}
	if len(g[0]) != 2 {
		http.Error(respW, "Unable to decode json", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(g[0][1])
	if err != nil {
		http.Error(respW, "Unable to decode json", http.StatusBadRequest)
		return
	}

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

// Implementation of the POST operation on product
// Adds to databse
func (p *Products) postProduct(respW http.ResponseWriter, req *http.Request) {
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

// Implementation of the GET operation on products
// Retrieves database
func (p *Products) getProducts(respW http.ResponseWriter, req *http.Request) {
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
