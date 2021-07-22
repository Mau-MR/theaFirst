package handlers

import (
	"github.com/Mau-MR/theaFirst/data"
	"github.com/Mau-MR/theaFirst/utils"
	"log"
	"net/http"
)

type Binnacles struct {
	l          *log.Logger
	BinnacleDB *data.CostumerDB
	validation *utils.Validation
}

func NewBinnacles(logger *log.Logger, BinnacleDB *data.CostumerDB, validation *utils.Validation) *Binnacles {
	return &Binnacles{
		l:          logger,
		BinnacleDB: BinnacleDB,
		validation: validation,
	}
}
func (bn *Binnacles) CreateBinnacle(rw http.ResponseWriter, r *http.Request) {
}
func (bn *Binnacles) SearchBinnacle(rw http.ResponseWriter, r *http.Request) {

}
func (bn *Binnacles) UpdateBinnacle(rw http.ResponseWriter, r *http.Request) {

}
