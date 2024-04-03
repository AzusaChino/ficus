package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongo(t *testing.T) {
	Init()
	defer func() {
		if err := Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := Client.Database("sample").Collection("movies")
	title := "back to the future"

	var result bson.M

	err := coll.FindOne(context.TODO(), bson.D{{Key: "title", Value: title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}

	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
