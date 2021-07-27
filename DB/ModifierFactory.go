package DB

import (
	"github.com/Mau-MR/theaFirst/connection"
)

//New creates a modifier based on the type of connection that is passed
func New(conn connection.Connection, db, collection string) Modifier {
	var modifier Modifier
	switch conn.(type) {
	case *connection.ElasticConnection:
		elastic := NewElasticModifier(conn.(*connection.ElasticConnection), db, collection)
		modifier = elastic
	case *connection.MongoConnection:
		mongo := &MongoModifier{}
		mongo.New(conn.(*connection.MongoConnection), db, collection)
		modifier = mongo
	}
	return modifier
}
