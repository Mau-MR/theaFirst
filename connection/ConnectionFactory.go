package connection

import "fmt"

//New creates a new connection based on the string that is passed and an error in case of failure
func New(connectionName string) (Connection, error) {
	var connection Connection
	switch connectionName {
	case "mongo":
		mongo := &MongoConnection{}
		if err := mongo.Connect(); err != nil {
			return nil, err
		}
		connection = mongo
	case "elasticsearch":
		elastic := &ElasticConnection{}
		if err := elastic.Connect(); err != nil {
			return nil, err
		}
		connection = elastic
	default:
		return nil, fmt.Errorf("Invalid connection name ")
	}
	return connection, nil
}
