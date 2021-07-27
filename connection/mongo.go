package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

//MongoConnection manages the necessary methods for matching the connection interface for mongo
type MongoConnection struct {
	Client *mongo.Client
}

//Connect create a new mongoConnection
func (mc *MongoConnection) Connect() error {
	mongoURI := os.Getenv("MONGOURI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	mc.Client = client
	return nil
}

//Close closes the mongo connection
func (mc *MongoConnection) Close() error {
	return mc.Client.Disconnect(context.Background())
}
