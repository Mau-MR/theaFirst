package data

import (
	"github.com/Mau-MR/theaFirst/DB"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BinnacleDB struct {
	mongo         *DB.MongoWrapper
	elasticSearch *DB.ElasticWrapper
	collection    string
}

//NewBinnacleDB returns a BinnacleDB for usage of the handler
func NewBinnacleDB(mongoClient *mongo.Client, elasticWrapper *DB.ElasticWrapper) *BinnacleDB {
	return &BinnacleDB{
		mongo:         DB.NewMongoWrapper("Thea", mongoClient),
		elasticSearch: elasticWrapper,
		collection:    "binnacles",
	}
}

//CreateBinnacle creates a new binnacle related with some costumer
func (bnc *BinnacleDB) CreateBinnacle(costumerID string) error {
	return nil
}

//CreateCell inserts a cell to the specified binnacle
func (bnc *BinnacleDB) CreateCell(binnacleID primitive.ObjectID, cell *BinnacleCell) error {
	return nil
}

//UpdateCell updates the specified cell of a binnacle
func (bnc *BinnacleDB) UpdateCell(cell *BinnacleCell) error {
	return nil
}
