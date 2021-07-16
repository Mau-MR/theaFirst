package main

import (
	"context"
	"github.com/Mau-MR/theaFirst/DB"
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
	//Env variables
	mongoURI := os.Getenv("MONGOURI")
	address := os.Getenv("ELASTICURI")
	username := os.Getenv("EUSER")
	password := os.Getenv("EPASSWORD")

	//The logger creation
	l := log.New(os.Stdout, "[Thea-API] ", log.LstdFlags)
	//DB Connections
	mongoWrapper, err := DB.NewMongoWrapper("Thea", mongoURI, l)
	if err != nil {
		l.Fatal("Unable to connect to MongoDB")
	}
	defer mongoWrapper.Disconnect(context.TODO())
	elasticSearchWrapper, err := DB.NewElasticWrapper(address, username, password, l)
	if err != nil {
		l.Fatal("Unable to connect to ElasticSearch: ", err)
	}

	//validator for every request
	validation := utils.NewValidation()

	//handlers
	costumers := handlers.NewCostumers(l, mongoWrapper, elasticSearchWrapper, validation)

	//Routes configuration
	mux := mux.NewRouter()
	//Post router
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/costumers", costumers.CreateCostumer)
	//Get router
	//Update router

	//server related configuration
	server := http.Server{
		//TODO: GET THE PORT FROM ENVIRONMENT
		Addr:         "localhost:8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  100 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 8080")
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
