package data

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

//Costumer The costumer struct has the general aspects of a client
type Costumer struct {
	Name         string             `json:"name" validate:"required" bson:"name"`
	Phone        string             `json:"phone" validate:"required" bson:"phone"` //TODO: CREATE A VALIDATION FOR THE PHONE NUMBER
	Organization string             `json:"organization" bson:"organization"`
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}

//NewCostumer  creates a new costumer structure an returns its pointer
func NewCostumer(name, phone, organization string) *Costumer {
	return &Costumer{
		Name:         name,
		Phone:        phone,
		Organization: organization,
	}
}
func (c *Costumer) ConvertToJSON(jsonString []byte) (interface{}, error) {
	costumer := &Costumer{}
	err := json.Unmarshal(jsonString, costumer)
	return costumer, err
}

//Standardize converts all fields of the struct to lowercase
func (c *Costumer) Standardize() {
	c.Name = strings.ToLower(c.Name)
	c.Organization = strings.ToLower(c.Organization)
}
