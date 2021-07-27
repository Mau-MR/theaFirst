package handlers

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data"
)

type CostumerDB struct {
	mongoDB       *DB.Modifier
	elasticSearch *DB.Modifier
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
func (c *CostumerDB) CreateCostumer(costumer *data.Costumer) error {
	return nil
}

func (c *CostumerDB) SearchCostumer(costumer *data.Costumer) ([]*data.Costumer, error) {
	return nil, nil
}
