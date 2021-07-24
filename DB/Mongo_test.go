package DB

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
)

//TODO CHECK FOR THIS TEST THAT THEY ARE DOING WHAT IS INTENDED TO BE
func TestNewMongoClient(t *testing.T) {
	mongoURI := os.Getenv("MONGOURI")
	_, err := NewMongoClient(mongoURI)
	assert.Nil(t, err)
}

func TestNewMongoWrapper(t *testing.T) {
	mongoURI := os.Getenv("MONGOURI")
	client, err := NewMongoClient(mongoURI)
	assert.Nil(t, err, "Error should not be nil")
	assert.NotNil(t, client, "Mongo client should not be nil")
	mongoWrapper := NewMongoWrapper("Test", client)
	assert.NotNil(t, mongoWrapper, "Mongo wrapper should not be nil")
}

func mongoWrapperTest(t *testing.T) *MongoWrapper {
	mongoURI := os.Getenv("MONGOURI")
	client, err := NewMongoClient(mongoURI)
	assert.Nil(t, err, "Error should not be nil")
	assert.NotNil(t, client, "Mongo client should not be nil")
	mongoWrapper := NewMongoWrapper("Test", client)
	assert.NotNil(t, mongoWrapper, "Mongo wrapper should not be nil")
	return mongoWrapper
}

func TestMongoWrapper_InsertStructTo(t *testing.T) {
	mongoWrapper := mongoWrapperTest(t)
	type Test struct {
		Name string `json:"name"`
	}
	test := Test{
		Name: "Mauricio",
	}
	res, err := mongoWrapper.Transaction(func(session mongo.SessionContext) (interface{}, error) {
		res, err := mongoWrapper.InsertStructTo("test", &session, test)
		if err != nil {
			return nil, err
		}
		return mongoWrapper.DeleteDocumentByID("test", res.InsertedID.(primitive.ObjectID), &session)
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
	//TODO: ADD  CODE HERE TO DELETE THE INSERTED DOCUMENT
}
/*
func TestMongoWrapper_SearchByFieldOn(t *testing.T) {
	mongoWrapper := mongoWrapperTest(t)
	costumer := &data.Costumer{Name: "Mauricio"}
	binnacle := &data.Binnacle{}
	binnacleReq := &data.PostBinnacleCellReq{
		Cell: data.BinnacleCell{
			NumberLashStart: 56,
			NumberLashEnd:   120,
		},
	}
	callback := func(ctx mongo.SessionContext) (interface{}, error) {
		costumerID, err := mongoWrapper.InsertStructTo("test", &ctx, costumer)
		if err != nil {
			return nil, err
		}
		binnacle.CostumerID = costumerID.InsertedID.(primitive.ObjectID)
		binnacleRes, err := mongoWrapper.InsertStructTo("test", &ctx, binnacle)
		if err != nil {
			return nil, err
		}
		binnacleReq.BinnacleID = binnacleRes.InsertedID.(primitive.ObjectID).Hex()
		_, err = mongoWrapper.InsertStructTo("test", &ctx, binnacleReq)
		if err != nil {
			return nil, err
		}
		return costumerID, err
	}
	res, err := mongoWrapper.Transaction(callback)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	costumerID := res.(mongo.InsertOneResult).InsertedID.(primitive.ObjectID)
	result := mongoWrapper.SearchByFieldOn("test", "clientID", costumerID)
	resultBinnacle := &data.Binnacle{}
	err = result.Decode(resultBinnacle)
}
func TestMongoWrapper_UpdateDocumentOn(t *testing.T) {
	mongoWrapper := mongoWrapperTest(t)
	newValues := &map[string]string{
		"name": "otherName here", //cannot update immutable id that's why id wasn't put on this test
	}
	type Test struct {
		Name string `json:"name"`
	}
	test := Test{Name: "Pedro"}
	res, err := mongoWrapper.Transaction(func(session mongo.SessionContext) (interface{}, error) {
		res, err := mongoWrapper.InsertStructTo("test", &session, test)
		if err != nil {
			return nil, err
		}
		_, err = mongoWrapper.UpdateDocumentOn("test", "name", "Pedro", &session, newValues)
		if err != nil {
			return nil, err
		}
		return mongoWrapper.DeleteDocumentByID("test", res.InsertedID.(primitive.ObjectID), &session)
	})
	assert.Nil(t, err)
	assert.NotNil(t, res)
}


func TestMongoWrapper_UpdateDocumentByID(t *testing.T) {
	type Test struct {
		Name string             `bson:"name"`
		ID   primitive.ObjectID `bson:"_id"`
	}
	test := Test{
		Name: "Mauricio",
	}
	newValues := &map[string]string{
		"name": "otherName here", //cannot update immutable id, that's why id wasn't put on this test
	}
	mongoWrapper := mongoWrapperTest(t)
	res, err := mongoWrapper.Transaction(func(session mongo.SessionContext) (interface{}, error) {
		return mongoWrapper.InsertStructTo("test", &session, test)
	})
	assert.NotNil(t, res)

	_, err = mongoWrapper.Transaction(func(session mongo.SessionContext) (interface{}, error) {
		resNew, err := mongoWrapper.SearchByFieldOn("test", "name", "Mauricio", test)
		_, err = mongoWrapper.UpdateDocumentByID("test", resNew., &session, newValues)
		if err != nil {
			return nil, err
		}
	return mongoWrapper.DeleteDocumentByID("test", resNew.(Test).ID, &session)
	})
	assert.Nil(t, err)
}

*/
