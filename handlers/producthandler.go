package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bal3000/go-products/infrastructure"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	dataStore infrastructure.DataStore
}

func NewProductHandler(ds infrastructure.DataStore) ProductHandler {
	return ProductHandler{dataStore: ds}
}

func (ph ProductHandler) GetPackSizes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	productId, ok := vars["productId"]

	if !ok {
		http.Error(w, "please provide a product id", http.StatusBadRequest)
	}

	sizes, ok, err := ph.dataStore.PackSizesForProduct(productId)
	if err != nil {
		log.Printf("error occured getting product from datastore, %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if !ok {
		http.Error(w, fmt.Sprintf("product %s not found", productId), http.StatusNotFound)
	}

	packSizes := struct {
		Sizes []int `json:"packSizes"`
	}{
		Sizes: sizes,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(packSizes); err != nil {
		log.Printf("failed to send sizes back to client, %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
