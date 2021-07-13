package main

import (
	"context"
	"crypto/tls"
	"github.com/Mau-MR/theaFirst/handlers"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func ElasticSearchClient(l *log.Logger) *elasticsearch.Client {
	//The retrieve of the credentials
	address := os.Getenv("ELASTICURI")
	username := os.Getenv("EUSER")
	password := os.Getenv("EPASSWORD")
	//The configuration  of the client
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		l.Fatal("Unable to connect to elasticsearch")
	}
	l.Println("Successfully connected to elasticsearch")
	return es
}
func MongoClient(l *log.Logger) (*mongo.Client, *context.Context) {
	mongoURI := os.Getenv("MONGOURI")
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		l.Fatal("Error assigning the URI")
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mongoClient.Connect(ctx)
	if err != nil {
		l.Fatal("Error Connecting to MongoDB")
	}
	l.Println("Mongo Successful connection")
	return mongoClient, &ctx
}

func main() {
	l := log.New(os.Stdout, "[Thea-API] ", log.LstdFlags)
	//elasticSearchClient, _ := elasticsearch.NewDefaultClient()
	mongoClient, ctx := MongoClient(l)
	defer mongoClient.Disconnect(*ctx)
	elasticSearchClient := ElasticSearchClient(l)
	costumers := handlers.NewCostumers(l,mongoClient,elasticSearchClient)
	mux:= mux.NewRouter()
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/costumers", costumers.CreateCostumer)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
		ErrorLog: l,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		IdleTimeout: 100*time.Second,
	}
	go  func(){
		l.Println("Starting server on port 8080")
		if err:= server.ListenAndServe(); err !=nil{
			l.Fatal("Error starting the server: ", err)
		}
	}()
	//get sigterm or interrupt to gracefully end the server
	c:= make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt)
	signal.Notify(c,os.Kill)
	//Block until signal is received
	sig:= <-c
	l.Println("Got signal", sig)
	//shutdonw the server and waiting 30 seconds for current operations to complete
	*ctx,_ = context.WithTimeout(context.Background(),30*time.Second)
	if err := server.Shutdown(*ctx); err!=nil{
		l.Fatal("Error shutting down the server",err)
	}
}
