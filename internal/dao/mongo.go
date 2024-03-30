package dao

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Init() {
	var err error
	err = godotenv.Load("local-test.env")
	if err != nil {
		log.Fatal(err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("no mongodb found")
	}

	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Panic(err)
	}

}
