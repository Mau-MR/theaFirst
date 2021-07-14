package DB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"time"
)

type MongoWrapper struct {
	client *mongo.Client
	l      *log.Logger
}
type callback func(mongo.SessionContext) (interface{}, error)

//NewMongoWrapper gets the mongo URI for a db and returns a MongoWraper with a client inside and a error in case of the failure of the connection
func NewMongoWrapper(cred string, logger *log.Logger) (*MongoWrapper, error) {
	mongoClient, err := mongoClient(cred)
	if err != nil {
		return nil, err
	}
	return &MongoWrapper{
		client: mongoClient,
		l:      logger,
	}, nil
}
func (m *MongoWrapper) Disconnect(context context.Context) error {
	m.l.Println("Disconnecting client from mongoDB")
	return m.client.Disconnect(context)
}

//TODO: CHECK FOR AUTHENTICATION SINCE THE FIRTS PETITION ALSO MAKES AUTHENTICATION AND IT IS SLOW ON THE FIRST REQUEST TO THE DB
//TODO: SEE IF THERE IS A WAY TO PRINT ERRORS WITHOUT BLOCKING THE PROGRAM WITH LOGGER
//mongoClient creates a client based on the URI that is passed and returns the client or a error in case of failure
func mongoClient(mongoURI string) (*mongo.Client, error) {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	return mongoClient, nil
}

//Transactions makes an ACID transaction for MongoDB
func (mw *MongoWrapper) Transaction(callback callback) error {
	wc := writeconcern.New(writeconcern.WMajority())
	rc := readconcern.Snapshot()
	txnOpts := options.Transaction().SetWriteConcern(wc).SetReadConcern(rc)
	session,err := mw.client.StartSession()
	if err !=nil{
		mw.l.Println("Unable to start the session")
		return err
	}
	defer session.EndSession(context.Background())
	//TODO: CHECK WHAT RETURNS THIS EXPRESION
	mw.l.Println("Starting transaction...")
	_,err = session.WithTransaction(context.Background(),callback,txnOpts)
	if err!=nil{
		mw.l.Println("Transaction failed")
		return err
	}
	mw.l.Println("Successful transaction!")
	return nil
}
//InsertStructTo inserts and struct on the specified db and collection returns the response of the insertion and an error in case of failure
func (mw *MongoWrapper) InsertStructTo(i interface{},db, collection string)(*mongo.InsertOneResult,error){
	return mw.client.Database(db).Collection(collection).InsertOne(context.Background(),i)
}
