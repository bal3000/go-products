package main

import (
	"log"

	"github.com/bal3000/go-products/app"
	"github.com/bal3000/go-products/infrastructure"
)

func main() {
	ds, err := infrastructure.NewJsonDatatore("./packsizes.json")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.NewServer(ds)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to startup server: %v", err)
	}
}
