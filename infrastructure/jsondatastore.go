package infrastructure

import (
	"embed"
	"encoding/json"
	"errors"
	"log"

	"github.com/bal3000/go-products/models"
)

type JsonDataStore struct {
	products map[string][]int
}

func NewJsonDatatore(fs embed.FS) (*JsonDataStore, error) {
	log.Println("loading json file")
	file, err := fs.Open("packsizes.json")
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

	return &JsonDataStore{products: pm}, nil
}

func (ds *JsonDataStore) PackSizesForProduct(productId string) ([]int, bool, error) {
	if ds == nil {
		return nil, false, errors.New("the data store has not been initialized")
	}

	sizes, ok := ds.products[productId]
	return sizes, ok, nil
}
