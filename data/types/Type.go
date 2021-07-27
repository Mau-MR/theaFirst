package types

import "go.mongodb.org/mongo-driver/bson/primitive"

//Type is the type that represent every type that is going to be inserted on the DB or the main  structures of the program
type Type interface {
	//String Used Only for formatting the prints
	String() string
	SearchTerm() string
	//SearchFields used for mongo searching by specific fields
	SearchFields() *map[string]string
	//SearchID
	PrimitiveIDs() *map[string]primitive.ObjectID
	Clone() *Type
	EmptyClone() *Type
	ID() primitive.ObjectID
	StringID() string
}
