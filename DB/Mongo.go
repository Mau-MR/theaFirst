package DB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"os"
	"time"
)

type MongoWrapper struct {
	client *mongo.Client
	l      *log.Logger
	db     *mongo.Database
}
type callback func(mongo.SessionContext) (interface{}, error)

//NewMongoWrapper gets the mongo URI for a db and returns a MongoWrapper with a client inside and a error in case of the failure of the connection
func NewMongoWrapper(db string, mongoClient *mongo.Client) *MongoWrapper {
	return &MongoWrapper{
		client: mongoClient,
		l:      log.New(os.Stdout, fmt.Sprintf("[%s-DB] ", db), log.LstdFlags),
		db:     mongoClient.Database(db),
	}
}

//NewMongoClient creates a client based on the URI that is passed and returns the client or a error in case of failure
func NewMongoClient(mongoURI string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	return client, err
}

//Transaction makes an ACID transaction for MongoDB
func (mw *MongoWrapper) Transaction(callback callback) (interface{}, error) {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	session, err := mw.client.StartSession()
	if err != nil {
		mw.l.Println("Unable to start the session")
		return nil, err
	}
	defer session.EndSession(context.Background())

	//TODO: CHECK WHAT RETURNS THIS EXPRESSION
	mw.l.Println("Starting transaction...")
	res, err := session.WithTransaction(context.Background(), callback, txnOpts)
	if err != nil {
		mw.l.Println("Transaction failed")
		return nil, err
	}
	mw.l.Println("Successful transaction!")
	return res, nil
}

//InsertStructTo inserts and struct on the specified db and collection returns the response of the insertion and an error in case of failure
func (mw *MongoWrapper) InsertStructTo(collection string, ctx *mongo.SessionContext, i interface{}) (*mongo.InsertOneResult, error) {
	return mw.db.Collection(collection).InsertOne(*ctx, i)
}

//TODO: CHECK A WAY TO OPTIMIZE with pointers

// SearchByFieldOn gets the collection key and value that is going to search an return the document in case of existence
func (mw *MongoWrapper) SearchByFieldOn(collection, key, value string, i interface{}) (interface{}, error) {
	if err := mw.db.Collection(collection).FindOne(context.Background(), bson.D{{Key: key, Value: value}}).Decode(&i); err != nil {
		return nil, err //means it could not be found
	}
	return i, nil
}
func (mw *MongoWrapper) SearchByID(collection string, ID primitive.ObjectID) *mongo.SingleResult {
	return mw.db.Collection(collection).FindOne(context.Background(), bson.D{{Key: "_id", Value: ID}})
}

//UpdateDocumentOn receives the search criteria with searchField and searchFieldValue and updates the document with the given newFieldsAndValues of type map[field]value
//NOTE: UpdateDocumentOn doesnt work with _id field, since we create another method UpdateDocumentByID to do this and handler better the type primitive.ObjectID
func (mw *MongoWrapper) UpdateDocumentOn(collection, searchField, searchFieldValue string, ctx *mongo.SessionContext, newFieldsAndValues *map[string]string) (interface{}, error) {
	changes := mw.convertMapToBsonD(newFieldsAndValues)
	mw.l.Println("Document changes: ", changes)
	return mw.db.Collection(collection).UpdateOne(
		*ctx,
		bson.M{searchField: searchFieldValue},
		bson.D{{
			"$set",
			&changes,
		}},
	)
}
func (mw *MongoWrapper) UpdateDocumentByID(collection string, id primitive.ObjectID, ctx *mongo.SessionContext, newFieldsAndValues *map[string]string) (interface{}, error) {
	changes := mw.convertMapToBsonD(newFieldsAndValues)

	return mw.db.Collection(collection).UpdateOne(
		*ctx,
		bson.M{"_id": id},
		bson.D{{
			"$set",
			&changes,
		}},
	)
}
func (mw *MongoWrapper) DeleteDocumentByID(collection string, id primitive.ObjectID, ctx *mongo.SessionContext) (interface{}, error) {
	return mw.db.Collection(collection).DeleteOne(*ctx, bson.M{"_id": id})
}

func (mw *MongoWrapper) convertMapToBsonD(newFieldsAndValues *map[string]string) *bson.D {
	var changes bson.D
	for k, v := range *newFieldsAndValues {
		changes = append(changes, bson.E{Key: k, Value: v}) //appending the fields that are going to be changed to the document
	}
	return &changes
}
