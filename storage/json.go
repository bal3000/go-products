package storage

import (
	"os"
	"encoding/json"
	"errors"
	"log"

	"github.com/bal3000/go-products/models"
)

type JSON struct {
	products map[string][]int
}

func NewJSON(fs string) (*JSON, error) {
	log.Println("loading json file")
	file, err := os.Open(fs)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	products := &models.ProductCollection{}
	err = json.NewDecoder(file).Decode(products)
	if err != nil {
		return nil, err
	}
	log.Println("loaded json file")

	pm := make(map[string][]int)
	for _, p := range products.Products {
		pm[p.ProductID] = p.PackSizes
	}

	return &JSON{products: pm}, nil
}

func (ds *JSON) PackSizesForProduct(productId string) ([]int, bool, error) {
	if ds == nil {
		return nil, false, errors.New("the data store has not been initialized")
	}

	sizes, ok := ds.products[productId]
	return sizes, ok, nil
}
