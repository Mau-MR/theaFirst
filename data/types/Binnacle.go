package types

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Binnacle struct {
	CostumerID string         `json:"costumerID,omitempty" bson:"costumerID,omitempty" validate:"required"`
	Records    []BinnacleCell `json:"records" bson:"records"`
	BinnacleID string         `json:"binnacleID,omitempty" bson:"binnacleID,omitempty"`
}

func (b *Binnacle) FromJSON(message json.RawMessage) error {
	err := json.Unmarshal(message, b)
	return err
}

func (b *Binnacle) SetID(id string) {
	b.BinnacleID = id
}

func (b *Binnacle) String() string {
	//TODO
	return ""
}

func (b *Binnacle) SearchTerm() string {
	return b.CostumerID
}

func (b *Binnacle) SearchIDS() (*map[string]primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(b.CostumerID)
	if err != nil {
		return nil, err
	}
	m := &map[string]primitive.ObjectID{
		"_id": id, //Since we search on the costumer db we are searching for the id of that document
	}
	return m, nil
}
func (b *Binnacle) SearchFields() *map[string]string {
	m := &map[string]string{
		"costumerID": b.CostumerID,
	}
	return m
}

func (b *Binnacle) Clone() Type {
	return b
}

func (b *Binnacle) EmptyClone() Type {
	return &Binnacle{}
}

func (b *Binnacle) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(b.BinnacleID)
	return id, err
}

func (b *Binnacle) StringID() string {
	return b.BinnacleID
}
