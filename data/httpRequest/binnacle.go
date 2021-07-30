package httpRequest

import (
	"encoding/json"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InsertBinnacleCell struct {
	BinnacleID string             `bson:"_id" json:"binnacleID" validate:"required"`
	Cell       types.BinnacleCell `json:"record" validate:"required"`
}

func (ib *InsertBinnacleCell) FromJSON(message json.RawMessage) error {
	return nil
}

func (ib *InsertBinnacleCell) SetID(id string) {
	ib.BinnacleID = id
}

func (ib *InsertBinnacleCell) String() string {
	return ""
}

func (ib *InsertBinnacleCell) SearchTerm() string {
	return ""
}

func (ib *InsertBinnacleCell) SearchFields() *map[string]string {
	m := &map[string]string{
		"binnacleID": ib.BinnacleID,
	}
	return m
}

func (ib *InsertBinnacleCell) SearchIDS() (*map[string]primitive.ObjectID, error) {
	//TODO
	return nil, nil
}

func (ib *InsertBinnacleCell) Clone() types.Type {
	return ib
}

func (ib *InsertBinnacleCell) EmptyClone() types.Type {
	return &ib.Cell
}

func (ib *InsertBinnacleCell) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(ib.BinnacleID)
	return id, err
}

func (ib *InsertBinnacleCell) StringID() string {
	return ib.BinnacleID
}
