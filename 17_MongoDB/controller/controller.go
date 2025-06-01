package controller

import (
	"context"
	"log"

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

	// context.TODO() / context.Background()
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(colName)

}
