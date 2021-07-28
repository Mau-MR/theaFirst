package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/handlers"
	"github.com/Mau-MR/theaFirst/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := flag.Int("port", 0, "The server port")
	flag.Parse()
	//The logger creation
	l := log.New(os.Stdout, "[Keybons-System] ", log.LstdFlags)
	//validator for every request
	validation := utils.NewValidation()
	mongoClient, elasticClient := CreateConnections(l)
	defer mongoClient.Close() //TODO: handle err
	//handlers
	costumers := handlers.NewCostumers(l, mongoClient, elasticClient, validation)
	binnacles := handlers.NewBinnacles(l, mongoClient, elasticClient, validation)
	//Routes configuration
	mux := mux.NewRouter()
	//Post router
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/costumer", costumers.CreateCostumer)
	postRouter.HandleFunc("/binnacle/cell", binnacles.CreateCell)
	postRouter.HandleFunc("/binnacle", binnacles.CreateBinnacle)
	//Get router
	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/costumer", costumers.SearchCostumer)
	getRouter.HandleFunc("/binnacle", binnacles.SearchBinnacle)
	//Update router TODO: CREATE THE RELATED METHODS
	//server related configuration
	server := http.Server{
		Addr:         fmt.Sprintf("localhost:%d", *port),
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  100 * time.Second,
	}
	go func() {
		l.Println("Starting server on port ", *port)
		if err := server.ListenAndServe(); err != nil {
			l.Fatal("Error starting the server: ", err)
		}
	}()
	//TODO: CHECK IF THIS CODE IS TRULY EXECUTED
	//get sigterm or interrupt to gracefully end the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	//Block until signal is received
	sig := <-c
	l.Println("Got signal", sig)
	//shutdown the server and waiting 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		l.Fatal("Error shutting down the server", err)
	}
}

func CreateConnections(l *log.Logger) (connection.Connection, connection.Connection) {
	//DB Connections
	mongoClient, err := connection.New("mongo")
	if err != nil {
		l.Fatal("Unable to connect to MongoDB")
	}
	elasticClient, err := connection.New("elasticsearch")
	if err != nil {
		l.Fatal("Unable to connect to ElasticSearch: ", err)
	}
	return mongoClient, elasticClient
}
