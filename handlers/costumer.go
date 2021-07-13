package handlers

import (
	"github.com/Mau-MR/theaFirst/data"
	"github.com/elastic/go-elasticsearch/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Costumers struct{
	l *log.Logger
	CostumerDB *data.CostumerDB
}
func NewCostumers(logger *log.Logger, mongoDBClient * mongo.Client,esClient *elasticsearch.Client) *Costumers{
	return &Costumers{
		l: logger,
		CostumerDB: data.NewCostumerDB(mongoDBClient,esClient),
	}
}

func(*Costumers) CreateCostumer(rw http.ResponseWriter, r*http.Request){

}