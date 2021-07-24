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

//CreateBinnacle insert a new binnacle on the db
func (bdb *BinnacleDB) CreateBinnacle(binnacle *Binnacle) error {
	callback := func(ctx mongo.SessionContext) (interface{}, error) {
		res, err := bdb.mongo.InsertStructTo(bdb.collection, &ctx, binnacle)
		if err != nil {
			return nil, err
		}
		return res, err
	}
	_, err := bdb.mongo.Transaction(callback)
	return err
}

//SearchBinnacle returns the binnacle that match a userID
func (bdb *BinnacleDB) SearchBinnacle(costumerID string) (*Binnacle, error) {
	id, err := primitive.ObjectIDFromHex(costumerID)
	if err != nil {
		return nil, err
	}
	bdb.mongo.SearchByFieldOn(bdb.collection, "costumerID", id)
	return nil, nil
}

//CreateCell inserts a cell to the specified binnacle
func (bdb *BinnacleDB) CreateCell(binnacleID primitive.ObjectID, cell *BinnacleCell) error {
	return nil
}

//UpdateCell updates the specified cell of a binnacle
func (bdb *BinnacleDB) UpdateCell(cell *BinnacleCell) error {
	return nil
}
