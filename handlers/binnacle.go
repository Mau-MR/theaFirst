package handlers

import (
	"github.com/Mau-MR/theaFirst/DB"
	"github.com/Mau-MR/theaFirst/data"
	"github.com/Mau-MR/theaFirst/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type Binnacles struct {
	l          *log.Logger
	BinnacleDB *data.BinnacleDB
	validation *utils.Validation
}

func NewBinnacles(logger *log.Logger, mongoClient *mongo.Client, elasticWrapper *DB.ElasticWrapper, validation *utils.Validation) *Binnacles {
	return &Binnacles{
		l:          logger,
		BinnacleDB: data.NewBinnacleDB(mongoClient, elasticWrapper),
		validation: validation,
	}
}
func (bs *Binnacles) CreateBinnacle(rw http.ResponseWriter, r *http.Request) {
	pBinnacleReq := &data.PostBinnacleReq{}
	err := utils.ParseRequest(pBinnacleReq, r.Body, rw)
	if err != nil {
		bs.l.Println("Error parsing account", err)
		return
	}
	err = bs.validation.ValidateRequest(pBinnacleReq, rw)
	if err != nil {
		bs.l.Println("Missing fields or validation error for costumer", pBinnacleReq)
		return
	}
	costumerID, err := primitive.ObjectIDFromHex(pBinnacleReq.CostumerID)
	if err != nil {
		bs.l.Println("Unable to convert the id", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{
			Message: "Unable to parse the id",
		}, rw)
		return
	}
	binnacle := &data.Binnacle{
		CostumerID: costumerID,
	}
	err = bs.BinnacleDB.CreateBinnacle(binnacle)
	if err != nil {
		bs.l.Println("Unable to create Binnacle", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{
			Message: "Unable to create the binnacle",
		}, rw)
		rw.WriteHeader(http.StatusOK)
		utils.ToJSON(data.NewSuccessfulRequest(), rw)
		bs.l.Println("Successfully created Binnacle")
		return
	}
}

//SearchBinnacle receives the data.GetBinnacleReq
func (bs *Binnacles) SearchBinnacle(rw http.ResponseWriter, r *http.Request) {
	binnacleReq := &data.GetBinnacleReq{}
	err := utils.ParseRequest(binnacleReq, r.Body, rw)
	if err != nil {
		bs.l.Println("Error parsing binnacleReq")
		return
	}
	err = bs.validation.ValidateRequest(binnacleReq, rw)
	if err != nil {
		bs.l.Println("Missing fields or validation error for binnacleReq", binnacleReq)
		return
	}
	binnacle, err := bs.BinnacleDB.SearchBinnacle(binnacleReq.ClientID)
	if err != nil {
		bs.l.Println("Unable to search for binnacle")
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{Message: "Unable to search binnacle on DB"}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	utils.ToJSON(binnacle, rw) //todo handler this error
	bs.l.Println("Successfully searched binnacle")
	return
}

func (bs *Binnacles) UpdateBinnacle(rw http.ResponseWriter, r *http.Request) {

}
func (bs *Binnacles) CreateCell(rw http.ResponseWriter, r *http.Request) {

}
