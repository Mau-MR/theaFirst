package handlers

import (
	"fmt"
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BinnacleDB struct {
	mongoBinnacles DB.Modifier
	mongoCostumers DB.Modifier
}

//NewBinnacleDB returns a BinnacleDB for usage of the handler
func NewBinnacleDB(mongoConnection connection.Connection, elasticConnection connection.Connection) *BinnacleDB {
	binnacles := DB.New(mongoConnection, "Thea", "binnacles")
	costumers := DB.New(mongoConnection, "Thea", "costumers")

	return &BinnacleDB{
		mongoBinnacles: binnacles,
		mongoCostumers: costumers,
	}
}

//CreateBinnacle insert a new binnacle on the db
func (bdb *BinnacleDB) CreateBinnacle(binnacle *types.Binnacle) error {
	_, err := bdb.mongoCostumers.SearchID(binnacle)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Costumer related to this binnacle doesnt exist")
	}
	err = bdb.mongoBinnacles.Insert(binnacle)
	if err != nil {
		return err
	}
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
