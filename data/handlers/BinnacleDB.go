package handlers

import (
	"context"
	"fmt"
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/httpRequest"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type BinnacleDB struct {
	mongoBinnacles DB.Modifier
	mongoCostumers DB.Modifier
	conn           *connection.MongoConnection
}

//NewBinnacleDB returns a BinnacleDB for usage of the handler
func NewBinnacleDB(mongoConnection connection.Connection, elasticConnection connection.Connection) *BinnacleDB {
	binnacles := DB.New(mongoConnection, "Thea", "binnacles")
	costumers := DB.New(mongoConnection, "Thea", "costumers")

	return &BinnacleDB{
		mongoBinnacles: binnacles,
		mongoCostumers: costumers,
		conn:           mongoConnection.(*connection.MongoConnection),
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
func (bdb *BinnacleDB) CreateCell(cell *httpRequest.InsertBinnacleCell) error {
	//Creating the cellID for the register
	cell.Cell.CellID = primitive.NewObjectID().Hex()
	err := bdb.mongoBinnacles.Push(cell)
	return err
}

//UpdateCell updates the specified cell of a binnacle
func (bdb *BinnacleDB) UpdateCell(cell *types.BinnacleCell) error {
	return nil
}

//HardCoded to quickly deliver
func (bdb *BinnacleDB) ValidateCell(cell *types.BinnacleCell) error {
	prod := &types.Product{}
	err := bdb.conn.Client.Database("Thea").Collection("products").FindOne(context.Background(), bson.M{"ID": cell.LashID}).Decode(prod)
	if err != nil {
		return fmt.Errorf("Unable to find related product")
	}
	emp := &types.Employee{}
	err = bdb.conn.Client.Database("Thea").Collection("employees").FindOne(context.Background(), bson.M{"ID": cell.EmployeeID}).Decode(emp)
	if err != nil {
		return fmt.Errorf("Unable to find employee")
	}
	svc := &types.Service{}
	err = bdb.conn.Client.Database("Thea").Collection("products").FindOne(context.Background(), bson.M{"ID": cell.ServiceID}).Decode(svc)
	if err != nil {
		return fmt.Errorf("Unable to find service")
	}
	return nil
}
