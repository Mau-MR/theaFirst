package types

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BinnacleCell struct {
	CellID          string `json:"cellID" bson:"cellID"`
	Date            string `json:"date" bson:"date" validate:"required"`                       //TODO validate date
	NextAppointment string `json:"nextAppointment" bson:"nextAppointment" validate:"required"` //Validate date
	ServiceID       string `json:"serviceID" bson:"serviceID" validate:"required"`
	Mapping         []int  `json:"mapping" bson:"mapping" validate:"required"`
	LashID          string `json:"lashID" bson:"lashID" validate:"required"`
	NumberLashStart int    `json:"nLashStart" bson:"nLashStart" validate:"required"`
	NumberLashEnd   int    `json:"nLashEnd" bson:"nLashEnd" validate:"required"`
	EmployeeID      string `json:"employeeID" bson:"employeeID" validate:"required"`
	Observations    string `json:"observations" bson:"observations" validate:"required"`
}

func (bc *BinnacleCell) FromJSON(message json.RawMessage) error {
	err := json.Unmarshal(message, bc)
	return err
}

func (bc *BinnacleCell) SetID(id string) {
	bc.CellID = id
}

func (bc *BinnacleCell) String() string {
	return ""
}

func (bc *BinnacleCell) SearchTerm() string {
	return ""
}

func (bc *BinnacleCell) SearchIDS() (*map[string]primitive.ObjectID, error) {
	return nil, nil
}
func (bc *BinnacleCell) SearchFields() *map[string]string {
	return nil
}

func (bc *BinnacleCell) Clone() Type {
	return bc
}

func (bc *BinnacleCell) EmptyClone() Type {
	return &Binnacle{}
}

func (bc *BinnacleCell) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(bc.CellID)
	return id, err
}

func (bc *BinnacleCell) StringID() string {
	return bc.CellID
}
