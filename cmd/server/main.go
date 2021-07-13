package main

import (
	"context"
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func ElasticSearchClient(l *log.Logger) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "thea",
		Password: "thea88!",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10*time.Second,
			DialContext:           (&net.Dialer{Timeout: 10*time.Second}).DialContext,
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

	res,err:= elasticSearchClient.Info()
	if err!=nil{
		l.Fatal(err)
	}
	l.Println(res)
}
