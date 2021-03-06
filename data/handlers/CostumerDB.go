package handlers

import (
	"fmt"
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/types"
)

type CostumerDB struct {
	mongoDB       DB.Modifier
	elasticSearch DB.Modifier
}

func NewCostumerDB(mongoConnection connection.Connection, elasticConnection connection.Connection) *CostumerDB {
	mongoModifier := DB.New(mongoConnection, "Thea", "costumers")
	elasticModifier := DB.New(elasticConnection, "Thea", "costumers")
	return &CostumerDB{
		mongoDB:       mongoModifier,
		elasticSearch: elasticModifier,
	}
}

//CreateCostumer receives a costumer type, creates a mongo transaction and inserts the same costumer into the elasticsearch cluster, if and err  occurs it is returned and the transaction ends
func (c *CostumerDB) CreateCostumer(costumer types.Type) error {
	_, err := c.mongoDB.SearchFields(costumer)
	if err == nil { //If err is nil means that costumer was found
		return fmt.Errorf("Costumer already exists")
	}
	err = c.mongoDB.Insert(costumer)
	if err != nil {
		return fmt.Errorf("Unable to insert to primary db: %v", err)
	}
	err = c.elasticSearch.Insert(costumer)
	if err != nil {
		return fmt.Errorf("Unable to insert to secondary db: %v", err)
	}
	return err
}

func (c *CostumerDB) SearchCostumer(csReq types.Type) ([]types.Type, error) {
	fields, err := c.elasticSearch.SearchFields(csReq)
	if err != nil {
		return nil, err
	}
	return fields, nil
}
