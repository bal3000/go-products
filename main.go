package main

import (
	"embed"
	"log"

	"github.com/bal3000/go-products/app"
	"github.com/bal3000/go-products/storage"
)

//go:embed packsizes.json
var packsizes embed.FS

func main() {
	ds, err := storage.NewJSON(packsizes)
	if err != nil {
		log.Fatalln(err)
	}

	// username := os.Getenv("MONGO_USERNAME")
	// password := os.Getenv("MONGO_PASSWORD")

	// ds, closer, err := storage.NewMongo(fmt.Sprintf("mongodb+srv://%s:%s@baltest.4maj7.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", username, password))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer closer()

	srv := app.NewServer(ds)

	if err := srv.Run(); err != nil {
		log.Fatalf("failed to startup server: %v", err)
	}
}
