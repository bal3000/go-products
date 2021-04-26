package infrastructure

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bal3000/go-products/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBDataStore struct {
	dbClient *mongo.Client
}

func NewDBDataStore(connString string) (*DBDataStore, func(), error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, nil, err
	}

	return &DBDataStore{dbClient: client}, func() {
		log.Println("closing mongo db connection")
		client.Disconnect(ctx)
	}, nil
}

func (ds *DBDataStore) PackSizesForProduct(productId string) ([]int, bool, error) {
	if ds == nil {
		return nil, false, errors.New("the data store has not been initialized")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database := ds.dbClient.Database("go-products")
	packsCol := database.Collection("packs")
	c, err := packsCol.Find(ctx, bson.M{})
	if err != nil {
		return nil, false, err
	}

	var products []models.Product
	err = c.All(ctx, &products)
	if err != nil {
		return nil, false, err
	}

	for _, p := range products {
		if p.ProductID == productId {
			return p.PackSizes, true, nil
		}
	}
	return nil, false, nil
}
