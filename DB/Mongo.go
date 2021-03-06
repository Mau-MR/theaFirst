package DB

import (
	"context"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoModifier struct {
	client     *connection.MongoConnection
	Collection *mongo.Collection
}
type callback func(mongo.SessionContext) (interface{}, error)

//New gets the mongo URI for a db and returns a MongoModifier with a client inside and a error in case of the failure of the connection
func (mw *MongoModifier) New(connection *connection.MongoConnection, db, collection string) {
	mw.client = connection
	mw.Collection = mw.client.Client.Database(db).Collection(collection)
}
func (mw *MongoModifier) Push(data types.Type) error {
	//Hard coded for this time just to deliver this feature as quick as posible
	id, err := data.PrimitiveID()
	if err != nil {
		return err
	}
	cell := data.EmptyClone()
	change := bson.M{"$push": bson.M{"records": cell}}
	filter := bson.M{"_id": id}
	_, err = mw.Collection.UpdateOne(context.Background(), filter, change)
	if err != nil {
		return err
	}
	return nil
}

/*
//Transaction makes an ACID transaction for MongoDB
func (mw *MongoModifier) Transaction(callback callback) (interface{}, error) {
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
*/

//Insert inserts and struct on the specified db and Collection returns the response of the insertion and an error in case of failure
func (mw *MongoModifier) Insert(data types.Type) error {
	res, err := mw.Collection.InsertOne(context.Background(), data)
	id := res.InsertedID.(primitive.ObjectID).Hex()
	data.SetID(id)
	log.Println(data.StringID())
	return err
}

// SearchFields gets the Collection key and value that is going to search an return the document in case of existence
func (mw *MongoModifier) SearchFields(data types.Type) ([]types.Type, error) {
	fieldsValue := data.SearchFields()
	var doc bson.D
	for key, val := range *fieldsValue {
		doc = append(doc, bson.E{Key: key, Value: val})
	}
	newType := data.EmptyClone()
	err := mw.Collection.FindOne(context.Background(), doc).Decode(newType)
	if err != nil {
		return nil, err
	}
	return []types.Type{newType}, err
}
func (mw *MongoModifier) SearchID(data types.Type) (types.Type, error) {
	newType := data.EmptyClone()
	fieldsID, err := data.SearchIDS()
	if err != nil {
		return nil, err
	}
	var doc bson.D
	for field, id := range *fieldsID {
		log.Println(id)
		doc = append(doc, bson.E{Key: field, Value: id})
	}
	err = mw.Collection.FindOne(context.Background(), doc).Decode(newType)
	return newType, err
}

func (mw *MongoModifier) Update(data types.Type) error {
	id, err := data.PrimitiveID()
	if err != nil {
		return err
	}
	_, err = mw.Collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.D{{
			Key:   "$set",
			Value: data,
		}},
	)
	return err
}

func (mw *MongoModifier) Delete(data types.Type) error {
	id, err := data.PrimitiveID()
	if err != nil {
		return err
	}
	_, err = mw.Collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
