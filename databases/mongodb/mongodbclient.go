package mongodbclient

import (
	"context"
	"ecomjc-be/config"
	database "ecomjc-be/databases/interface"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbClient struct {
	client *mongo.Client
}

func NewMongoDBClient(conf *config.Config) database.Interface {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.MongoURI))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return &mongodbClient{
		client: client,
	}
}
