package data

import (
	"github.com/elastic/go-elasticsearch/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

type CostumerDB struct{
	mongoDB *mongo.Client
	elasticsearch *elasticsearch.Client
}
func NewCostumerDB(mongoClient *mongo.Client,elasticsearchClient *elasticsearch.Client) *CostumerDB{
	return &CostumerDB{
		mongoDB: mongoClient,
		elasticsearch: elasticsearchClient,
	}
}
