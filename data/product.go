package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

// GetProducts returns all products
func GetProducts() Products {
	return productList
}

// AddProduct adds a product to the list
func AddProduct(p *Product) {
	productList = append(productList, p)
}

var productList = Products{
	&Product{
		Name:        "Guitar",
		Description: "Furch Orange Guitar",
		Price:       2000,
		SKU:         "aaa000",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		Name:        "Drums",
		Description: "Yamaha Drums",
		Price:       1500,
		SKU:         "bbb000",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
