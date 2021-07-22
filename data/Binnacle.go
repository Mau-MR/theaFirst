package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Binnacle struct {
	ClientID   string         `json:"clientID" bson:"clientID"`
	Records    []BinnacleCell `json:"records" bson:"records"`
	BinnacleID string         `json:"binnacleID" bson:"binnacleID"`
}

type BinnacleCell struct {
	CellID          string              `json:"cellID" bson:"cellID"`
	Date            primitive.Timestamp `json:"date" bson:"date"`
	NextAppointment primitive.Timestamp `json:"nextAppointment" bson:"nextAppointment"`
	ServiceID       string              `json:"serviceID" bson:"serviceID"`
	Mapping         []int               `json:"mapping" bson:"mapping"`
	LashID          string              `json:"lashID" bson:"lashID"`
	NumberLashStart int                 `json:"nLashStart" bson:"nLashStart"`
	NumberLashEnd   int                 `json:"nLashEnd" bson:"nLashEnd"`
	EmployeeID      string              `json:"employeeID" bson:"employeeID"`
	Observations    string              `json:"observations" bson:"observations"`
}
