package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	r.HandleFunc("/api/products", handler.CalculatePacksToSend).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		log.Println("started server on port 8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("shutting down server")
	os.Exit(0)

	return nil
}
