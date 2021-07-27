package types

import (
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

//String formats the type for being printed with fmt

func (c *Costumer) String() string {
	//TODO IMPLEMENTATION
	return ""
}

func (c *Costumer) SearchTerm() string {
	//TODO
	return ""
}

func (c *Costumer) SearchFields() *map[string]string {
	//TODO
	m := make(map[string]string)
	return &m
}

func (c *Costumer) Clone() Type {
	//TODO
	return c
}

func (c *Costumer) EmptyClone() Type {
	return &Costumer{}
}

func (c *Costumer) PrimitiveID() (primitive.ObjectID, error) {
	id, err := primitive.ObjectIDFromHex(c.ID)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (c *Costumer) StringID() string {
	return c.ID
}
