package app

import (
	"net/http"
	"time"

	"github.com/bal3000/go-products/handlers"
	"github.com/bal3000/go-products/infrastructure"
	"github.com/gorilla/mux"
)

type Server struct {
	dataStore infrastructure.DataStore
}

func NewServer(ds infrastructure.DataStore) Server {
	return Server{dataStore: ds}
}

func (s Server) Run() error {
	handler := handlers.NewProductHandler(s.dataStore)

	r := mux.NewRouter()
	r.HandleFunc("/api/products/{productId}", handler.GetPackSizes).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	return nil
}
