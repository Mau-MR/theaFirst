package httpRequest

import (
	"encoding/json"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SearchCostumer Is the request definition for searching a specific costumer
type SearchCostumer struct {
	Name string `bson:"name" json:"name"`
	ID   string `bson:"name" json:"_id"`
}

func (sc *SearchCostumer) String() string {
	//todo
	return ""
}
func (sc *SearchCostumer) SetID(newID string) {
	sc.ID = newID
}
func (sc *SearchCostumer) FromJSON(m json.RawMessage) error {
	err := json.Unmarshal(m, sc)
	return err
}

func (sc *SearchCostumer) SearchTerm() string {
	//TODO
	return ""
}

func (sc *SearchCostumer) SearchFields() *map[string]string {
	m := &map[string]string{
		"name": sc.Name,
	}
	return m
}

func (sc *SearchCostumer) Clone() types.Type {
	return sc
}

func (sc *SearchCostumer) EmptyClone() types.Type {
	return &SearchCostumer{}
}

func (sc *SearchCostumer) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(sc.Name)
	return id, err
}

func (sc *SearchCostumer) StringID() string {
	//TODO
	return ""
}
