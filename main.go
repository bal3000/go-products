package main

import (
	"embed"
	"log"

	"github.com/bal3000/go-products/app"
	"github.com/bal3000/go-products/infrastructure"
)

//go:embed packsizes.json
var packsizes embed.FS

func main() {
	ds, err := infrastructure.NewJsonDatatore(packsizes)
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.NewServer(ds)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to startup server: %v", err)
	}
}
