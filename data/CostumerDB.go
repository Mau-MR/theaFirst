package data

import (
	"github.com/Mau-MR/theaFirst/DB"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CostumerDB struct {
	mongoDB       *DB.MongoWrapper
	elasticSearch *DB.ElasticWrapper
	collection    string
}

func NewCostumerDB(mongoWrapper *DB.MongoWrapper, elasticsearchClient *DB.ElasticWrapper, collection string) *CostumerDB {
	return &CostumerDB{
		mongoDB:       mongoWrapper,
		elasticSearch: elasticsearchClient,
		collection:    collection,
	}
}

//CreateCostumer receives a costumer type, creates a mongo transaction and inserts the same costumer into the elasticsearch cluster, if and err  occurs it is returned and the transaction ends
func (c *CostumerDB) CreateCostumer(costumer *Costumer) error {
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		res, err := c.mongoDB.InsertStructTo(costumer,&sessionContext, c.collection)
		if err != nil {
			return nil, err
		}
		_, err = c.elasticSearch.InsertStructTo(costumer, c.collection, res.InsertedID.(primitive.ObjectID).Hex())
		return nil, err
	}
	return c.mongoDB.Transaction(callback)
}
