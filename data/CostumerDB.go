package data

import (
	"github.com/Mau-MR/theaFirst/DB"
	"go.mongodb.org/mongo-driver/mongo"
)

type CostumerDB struct {
	mongoDB       *DB.MongoWrapper
	elasticSearch *DB.ElasticWrapper
	db            string
	collection    string
}

func NewCostumerDB(mongoWrapper *DB.MongoWrapper, elasticsearchClient *DB.ElasticWrapper, db, collection string) *CostumerDB {
	return &CostumerDB{
		mongoDB:       mongoWrapper,
		elasticSearch: elasticsearchClient,
		db:            db,
		collection:    collection,
	}
}

func (c *CostumerDB) CreateCostumer(costumer *Costumer) error {
	callback := func(mongo.SessionContext) (interface{}, error) {
		res, err := c.mongoDB.InsertStructTo(costumer, c.db, c.collection)
		if err != nil {
			return nil, err
		}
		//TODO: ADD HERE THE CREATION OF THE DOCUMENT FOR ELASTICSEARCH
		return res, err
	}
	return c.mongoDB.Transaction(callback)
}
