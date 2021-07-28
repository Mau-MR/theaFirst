package types

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	//TODO
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
