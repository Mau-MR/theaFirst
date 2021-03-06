package DB

import (
	"github.com/Mau-MR/theaFirst/data/types"
)

//Modifier is the type related to the modifications of the dbs
type Modifier interface {
	Insert(data types.Type) error
	//Update updates a type based on their PrimitiveID with the new struct data
	Update(data types.Type) error
	Push(data types.Type) error
	SearchFields(data types.Type) ([]types.Type, error)
	SearchID(data types.Type) (types.Type, error)
	//Delete deletes a document based on their PrimitiveID
	Delete(data types.Type) error

	/*
		TODO LATER
		InsertMany(dataArray *[]data.Type) error
		UpdateMany(dataArray *[]data.Type) error
		SearchMany(dataArray *[]data.Type) (*[]data.Type, error)
		DeleteMany(dataArray *[]data.Type) error
	*/
}
