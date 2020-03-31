package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

//Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON decodes a product structured in JSON into usable text
func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

// AddProduct appends product item to end of th list
func AddProduct(p *Product) {
	p.ID = getNextID()
	p.CreatedOn = time.Now().UTC().String()
	p.UpdatedOn = p.CreatedOn
	productList = append(productList, p)
}

// UpdateProduct updates
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.UpdatedOn = time.Now().UTC().String()
	productList[pos] = p
	return nil
}

// ErrProductNotFound struct defines error
var ErrProductNotFound = fmt.Errorf("Product not foun")

// returns product iof compatible id
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if id == p.ID {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// getNextID returns next ID for insertion at list
func getNextID() int {
	return (productList[len(productList)-1].ID) + 1
}

// Products defines an array o type &Product
type Products []*Product

// ToJSON mehtod encodes contents of databse to JSON
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns database
func GetProducts() Products {
	return productList
}

// Static Database
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          3,
		Name:        "MilkChocollate",
		Description: "Chocollate with milk",
		Price:       3.15,
		SKU:         "jns01",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
