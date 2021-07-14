package data

//Costumer The costumer struct has the general aspects of a client
type Costumer struct{
	Name  string `json:"name" validate:"required" bson:"name"`
	Phone string `json:"phone" validate:"required" bson:"phone"`
}
//NewCostumer  creates a new costumer structure an returns its pointer
func NewCostumer(name,phone string)*Costumer{
	return &Costumer{
		Name: name,
		Phone:phone,
	}
}