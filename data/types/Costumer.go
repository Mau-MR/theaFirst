package types

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
)

//Costumer The costumer struct has the general aspects of a client
type Costumer struct {
	Name         string `json:"name" validate:"required" bson:"name"`
	Phone        string `json:"phone" validate:"required" bson:"phone"` //TODO: CREATE A VALIDATION FOR THE PHONE NUMBER
	Organization string `json:"organization" bson:"organization"`
	ID           string `json:"_id,omitempty" bson:"_id,omitempty"`
}

//Standardize converts all fields of the struct to lowercase
func (c *Costumer) Standardize() {
	c.Name = strings.ToLower(c.Name)
	c.Organization = strings.ToLower(c.Organization)
}
func (c *Costumer) SetID(newID string) {
	c.ID = newID
}

//String formats the type for being printed with fmt

func (c *Costumer) String() string {
	//TODO IMPLEMENTATION
	return ""
}

func (c *Costumer) SearchTerm() string {
	//TODO
	return ""
}

func (c *Costumer) SearchIDS() (*map[string]primitive.ObjectID, error) {
	return nil, nil
}
func (c *Costumer) SearchFields() *map[string]string {
	m := &map[string]string{
		"phone": c.Phone,
	}
	return m
}

func (c *Costumer) Clone() Type {
	//TODO
	return c
}

func (c *Costumer) EmptyClone() Type {
	return &Costumer{}
}

func (c *Costumer) FromJSON(m json.RawMessage) error {
	err := json.Unmarshal(m, c)
	return err
}
func (c *Costumer) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(c.ID)
	return id, err
}

func (c *Costumer) StringID() string {
	return c.ID
}
