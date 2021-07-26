package DB

import (
	"github.com/Mau-MR/theaFirst/connection"
)

func New(conn connection.Connection, db, collection string) *Modifier {
	var modifier *Modifier
	switch conn.(type) {
	case *connection.ElasticConnection:

	case *connection.MongoConnection:
	}
	return modifier

}
