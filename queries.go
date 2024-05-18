package main

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Customer struct {
	Name         string `bson:"name"`
	EmailAddress string `bson:"email_address"`
}

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	defer client.Disconnect(context.TODO())
	database := client.Database("mydb")
	database.CreateCollection(context.TODO(), "customers")
	collection := database.Collection("customers")

	document := Customer{Name: "Charlie", EmailAddress: "charlie@gmail.com"}
	collection.InsertOne(context.TODO(), document)

	document = Customer{Name: "Bob", EmailAddress: "bob@gmail.com"}
	collection.InsertOne(context.TODO(), document)

	document = Customer{Name: "Alice", EmailAddress: "alice@outlook.com"}
	collection.InsertOne(context.TODO(), document)

	searchQuery := bson.D{{Key: "email_address", Value: "bob@gmail.com"}}
	var customer bson.M
	// collection.FindOne(context.TODO(), searchQuery).Decode(&customer)
	// retrieves the first result only
	// collection.Distinct(context.TODO(), "name", bson.D{}) retrieves all distinct names
	// searchQuery := bson.D{{Key: "email_address", Value: bson.D{{Key: "$regex", Value: "^bob@"}}}}
	// retrieves all documents with email addresses starting with "bob@"
	cursor, _ := collection.Find(context.TODO(), searchQuery)
	cursor.Next(context.TODO())
	cursor.Decode(&customer)
	fmt.Println(customer)
	jsonbytes, _ := json.Marshal(customer)
	fmt.Println(string(jsonbytes))
	// use _, hasField := customer["field"] to see if a field exists (hasField is a bool)
	fmt.Println(customer["name"])
	fmt.Println(customer["email_address"])

	searchQuery = bson.D{{Key: "email_address", Value: "alice@outlook.com"}}
	updatedField := bson.D{{Key: "email_address", Value: "alice@gmail.com"}}
	updater := bson.D{{Key: "$set", Value: updatedField}}
	collection.UpdateMany(context.TODO(), searchQuery, updater)

	searchQuery = bson.D{{Key: "email_address", Value: "charlie@gmail.com"}}
	collection.DeleteMany(context.TODO(), searchQuery)

	// use Value: -1 to sort in reverse (descending) order
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: 1}})
	cursor, _ = collection.Find(context.TODO(), bson.D{}, opts)
	var res []bson.M
	cursor.All(context.TODO(), &res)
	for _, document := range res {
		fmt.Println(document)
	}

	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	fmt.Println(count)
}
