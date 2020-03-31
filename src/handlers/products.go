package handlers

import (
	"BuildingMicroservicesWithGo_NicJackson/src/data"
	"log"
	"net/http"
)

//Products defines a handler for product
type Products struct {
	l *log.Logger
}

//NewProducts defines a function that returns a new Products object handler
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	//handle Get operation requests
	if req.Method == http.MethodGet {
		p.getProducts(respW, req)
		return
	}

	//handle Put operation requests

	//Catch all
	respW.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(respW http.ResponseWriter, req *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(respW)
	if err != nil {
		http.Error(respW, "Unable to marshal json", http.StatusInternalServerError)
	}
}
