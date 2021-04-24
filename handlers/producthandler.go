package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bal3000/go-products/infrastructure"
	"github.com/bal3000/go-products/models"
	"github.com/bal3000/go-products/packs"
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

func (ph ProductHandler) CalculatePacksToSend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	request := &models.PackRequest{}
	err := json.NewDecoder(r.Body).Decode(request)
	if err != nil {
		http.Error(w, "there was a problem with your request, please check your request body", http.StatusBadRequest)
	}

	sizes, ok, err := ph.dataStore.PackSizesForProduct(request.ProductID)
	if err != nil {
		log.Printf("error occured getting product from datastore, %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	if !ok {
		http.Error(w, fmt.Sprintf("product %s not found", request.ProductID), http.StatusNotFound)
	}

	packsToSend := packs.CalculatePackSizes(sizes, request.ItemsOrder)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(packsToSend); err != nil {
		log.Printf("failed to send total back to client, %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
