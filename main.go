package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prakashsingha/orderAPI/config"
	"github.com/prakashsingha/orderAPI/services"
	"github.com/prakashsingha/orderAPI/routers"

	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	log.Printf("Server started...")

	client, ctx := connectDB()
	defer client.Disconnect(ctx)

	router := routers.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

func connectDB() (*mongo.Client, context.Context) {
	conf := config.GetConfig()
	client, err := services.NewClient(conf)

	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	services.SetClient(client)

	return client, ctx
}
