package types

import "go.mongodb.org/mongo-driver/bson/primitive"

//Type is the type that represent every type that is going to be inserted on the DB or the main  structures of the program
type Type interface {
	//String Used Only for formatting the prints
	String() string
	SearchTerm() string
	//SearchFields used for mongo searching by specific fields
	SearchFields() *map[string]string
	Clone() Type
	EmptyClone() Type
	PrimitiveID() (primitive.ObjectID,error)
	StringID() string
}
