package DB

import (
	"github.com/Mau-MR/theaFirst/connection"
)

func New(conn connection.Connection, db, collection string) *Modifier {
	var modifier Modifier
	switch conn.(type) {
	case *connection.ElasticConnection:
		mongo := &MongoModifier{}
		mongo.New(conn, db, collection)
	case *connection.MongoConnection:
	}
	return modifier
}
