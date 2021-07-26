package DB

import (
	"github.com/Mau-MR/theaFirst/data"
)

//Modifier is the type related to the modifications of the dbs
type Modifier interface {
	Insert(data data.Type) error
	//Update updates a type based on their ID with the new struct data
	Update(data data.Type) error
	SearchFields(data data.Type) (*data.Type, error)
	SearchID(data data.Type) (*data.Type, error)
	//Delete deletes a document based on their ID
	Delete(data data.Type) error

	/**
	TODO LATER
	InsertMany(dataArray *[]data.Type) error
	UpdateMany(dataArray *[]data.Type) error
	SearchMany(dataArray *[]data.Type) (*[]data.Type, error)
	DeleteMany(dataArray *[]data.Type) error
	*/
}
