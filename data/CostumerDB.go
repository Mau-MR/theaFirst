package data

import (
	"errors"
	"github.com/Mau-MR/theaFirst/DB"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CostumerDB struct {
	mongoDB       *DB.MongoWrapper
	elasticSearch *DB.ElasticWrapper
	collection    string
}

func NewCostumerDB(mongoClient *mongo.Client, elasticsearchClient *DB.ElasticWrapper) *CostumerDB {
	return &CostumerDB{
		mongoDB:       DB.NewMongoWrapper("Thea", mongoClient),
		elasticSearch: elasticsearchClient,
		collection:    "costumers",
	}
}

//CreateCostumer receives a costumer type, creates a mongo transaction and inserts the same costumer into the elasticsearch cluster, if and err  occurs it is returned and the transaction ends
func (c *CostumerDB) CreateCostumer(costumer *Costumer) error {
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		_, err := c.mongoDB.SearchByFieldOn(c.collection, "phone", costumer.Phone, Costumer{})
		if err == nil { //means the user with that phone already exist
			return costumer, errors.New("user already exists")
		}
		res, err := c.mongoDB.InsertStructTo(c.collection, &sessionContext, costumer)
		if err != nil {
			return nil, err
		}
		_, err = c.elasticSearch.InsertStructTo(c.collection, res.InsertedID.(primitive.ObjectID).Hex(), costumer)
		return nil, err
	}
	_, err := c.mongoDB.Transaction(callback)
	return err
}

func (c *CostumerDB) SearchCostumer(costumer *Costumer) ([]*Costumer, error) {
	query := c.elasticSearch.BuildSearchQueryByFields(costumer.Name, []string{"name"})
	//getting the response
	rw, err := c.elasticSearch.SearchIn(c.collection, query)
	if err != nil {
		return nil, err
	}
	var costumers []*Costumer
	//getting the id of the documents that match
	for _, hit := range rw.Hits.Hits {
		id, err := primitive.ObjectIDFromHex(hit.ID)
		if err != nil {
			return nil, err
		}
		costumer := Costumer{}
		oneResult := c.mongoDB.SearchByID(c.collection, id)
		err = oneResult.Decode(&costumer)
		if err != nil {
			return costumers, err
		}
		costumers = append(costumers, &costumer)
	}
	return costumers, err
}
