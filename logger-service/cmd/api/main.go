package main

import (
	"context"
	"fmt"
	"log"
	"logger-service/cmd/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "8082"
	rpcPort  = "5001"
	mongoURL = "mongodb://admin:password@mongo:27017/logs"
	gRpcPort = "5001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	log.Println("Starting logger service ! ...")

	// connect mongo db
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected to mongo db !")

	client = mongoClient

	app := Config{
		Models: data.New(client),
	}

	// start wen server
	// create a context in order to disconnect db
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	if err := app.serve(); err != nil {
		log.Panic("Server error:", err)
	}

}

func (app *Config) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Printf("Server starting on port %s\n", webPort)
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("Server error: %v\n", err)
		return err
	}
	return nil
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	conn, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}
	log.Println("mongo db connected")
	return conn, nil
}
