package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/geekAshish/mongodb/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchList"

// MOST IMPORTANT
var collection *mongo.Collection


// connect to db
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// In version 1: context.TODO() / context.Background() 
	// version 2 just clientoption
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection = client.Database(dbName).Collection(colName)

}


// MONGODB helpers - file
// insert 1 record


func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update);

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result.ModifiedCount)
}