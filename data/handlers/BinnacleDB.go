package handlers

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BinnacleDB struct {
	mongo         DB.Modifier
	elasticSearch DB.Modifier
	collection    string
}

//NewBinnacleDB returns a BinnacleDB for usage of the handler
func NewBinnacleDB(mongoConnection connection.Connection, elasticConnection connection.Connection) *BinnacleDB {
	mongoModifier := DB.New(mongoConnection, "Thea", "costumers")
	elasticModifier := DB.New(elasticConnection, "Thea", "costumers")
	return &BinnacleDB{
		mongo:         mongoModifier,
		elasticSearch: elasticModifier,
		collection:    "binnacles",
	}
}

//CreateBinnacle insert a new binnacle on the db
func (bdb *BinnacleDB) CreateBinnacle(binnacle *types.Binnacle) error {
	return nil
}

//SearchBinnacle returns the binnacle that match a userID
func (bdb *BinnacleDB) SearchBinnacle(costumerID string) (*types.Binnacle, error) {
	return nil, nil
}

//CreateCell inserts a cell to the specified binnacle
func (bdb *BinnacleDB) CreateCell(binnacleID primitive.ObjectID, cell *types.BinnacleCell) error {
	return nil
}

//UpdateCell updates the specified cell of a binnacle
func (bdb *BinnacleDB) UpdateCell(cell *types.BinnacleCell) error {
	return nil
}
