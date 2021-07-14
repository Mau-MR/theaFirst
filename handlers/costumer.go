package handlers

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/data"
	"github.com/Mau-MR/theaFirst/utils"
	"log"
	"net/http"
)

type Costumers struct {
	l          *log.Logger
	CostumerDB *data.CostumerDB
	validation *utils.Validation
}

func NewCostumers(logger *log.Logger, mongoWrapper *DB.MongoWrapper, elasticSearchWrapper *DB.ElasticWrapper, validation *utils.Validation) *Costumers {
	//NOTE: FOR THIS TIME THIS IS GOING TO BE HARD CODED BUT IT CAN BE DYNAMICALLY PROVISIONED
	db := "Thea" //TODO: ADD HERE DINAMIC PROVISION of the db
	collection := "costumers"
	return &Costumers{
		l:          logger,
		CostumerDB: data.NewCostumerDB(mongoWrapper, elasticSearchWrapper, db, collection),
		validation: validation,
	}
}

func (c *Costumers) CreateCostumer(rw http.ResponseWriter, r *http.Request) {
	costumer := &data.Costumer{}
	err := utils.ParseRequest(costumer, r.Body, rw)
	if err != nil {
		c.l.Println("Error parsing account", err)
		return
	}
	err = c.validation.ValidateRequest(costumer, rw)
	if err != nil {
		c.l.Println("Missing fields or validation error for costumer", costumer)
		return
	}
	err = c.CostumerDB.CreateCostumer(costumer)
	if err != nil {
		c.l.Println("Unable to insert on dbs: ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{
			Message: "Unable to insert into DB",
		}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	utils.ToJSON(data.NewSuccessfulRequest(), rw)
	c.l.Println("Successfully created costumer")
	return
}
