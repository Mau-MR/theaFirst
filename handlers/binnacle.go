package handlers

import (
	"fmt"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data"
	"github.com/Mau-MR/theaFirst/data/handlers"
	"github.com/Mau-MR/theaFirst/data/httpRequest"
	"github.com/Mau-MR/theaFirst/data/types"
	"github.com/Mau-MR/theaFirst/utils"
	"log"
	"net/http"
	"time"
)

type Binnacles struct {
	l          *log.Logger
	BinnacleDB *handlers.BinnacleDB
	validation *utils.Validation
}

func NewBinnacles(logger *log.Logger, mongoClient connection.Connection, elasticWrapper connection.Connection, validation *utils.Validation) *Binnacles {
	return &Binnacles{
		l:          logger,
		BinnacleDB: handlers.NewBinnacleDB(mongoClient, elasticWrapper),
		validation: validation,
	}
}
func (bs *Binnacles) CreateBinnacle(rw http.ResponseWriter, r *http.Request) {
	binnacle := &types.Binnacle{}
	err := utils.ParseRequest(binnacle, r.Body, rw)
	if err != nil {
		bs.l.Println("Error parsing account", err)
		return
	}
	err = bs.validation.ValidateRequest(binnacle, rw)
	if err != nil {
		bs.l.Println("Missing fields or validation error for costumer", binnacle)
		return
	}
	err = bs.BinnacleDB.CreateBinnacle(binnacle)
	if err != nil {
		bs.l.Println("Unable to create Binnacle", err)
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{
			Message: "Unable to create the binnacle",
		}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	utils.ToJSON(data.NewSuccessfulRequest(), rw)
	bs.l.Println("Successfully created Binnacle")
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
	cell := &httpRequest.InsertBinnacleCell{}
	err := utils.ParseRequest(cell, r.Body, rw)
	if err != nil {
		bs.l.Println("Error parsing cell")
		return
	}
	err = bs.validation.ValidateRequest(cell, rw)
	if err != nil {
		bs.l.Println("Missing fields or validation error for cell", cell)
		return
	}
	err = bs.validateCell(&cell.Cell)
	if err != nil {
		bs.l.Println("invalid cell content")
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{Message: "Bad format for binnacle cell"}, rw)
		return
	}
	err = bs.BinnacleDB.CreateCell(cell)
	if err != nil {
		bs.l.Println("Unable to search for binnacle")
		rw.WriteHeader(http.StatusInternalServerError)
		utils.ToJSON(utils.GenericError{Message: "Unable to search binnacle on DB"}, rw)
		return
	}
	rw.WriteHeader(http.StatusOK)
	utils.ToJSON(data.NewSuccessfulRequest(), rw)
	bs.l.Println("Successfully added cell")
	return
}

//Hardcoded for quickly delivery
func (bs *Binnacles) validateCell(cell *types.BinnacleCell) error {
	//Validate dates
	_, err := validateTime(cell.Date)
	_, err = validateTime(cell.NextAppointment)
	if err != nil {
		return fmt.Errorf("Invalid date")
	}
	err = bs.BinnacleDB.ValidateCell(cell)
	return err
}
func validateTime(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date)
}
