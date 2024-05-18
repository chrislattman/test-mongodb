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
	res, _ := collection.InsertOne(context.TODO(), document)
	if res.InsertedID == nil {
		fmt.Println("insert Charlie failed")
	}

	document = Customer{Name: "Bob", EmailAddress: "bob@gmail.com"}
	res, _ = collection.InsertOne(context.TODO(), document)
	if res.InsertedID == nil {
		fmt.Println("insert Bob failed")
	}

	document = Customer{Name: "Alice", EmailAddress: "alice@outlook.com"}
	res, _ = collection.InsertOne(context.TODO(), document)
	if res.InsertedID == nil {
		fmt.Println("insert Alice failed")
	}

	data := []interface{} {
		Customer{Name: "Daniel", EmailAddress: "daniel@gmail.com"},
		Customer{Name: "Frank", EmailAddress: "frank@gmail.com"},
	}
	_, err := collection.InsertMany(context.TODO(), data)
	if err != nil {
		fmt.Println("insert Daniel and Frank failed")
	}

	// indexModel := mongo.IndexModel{Keys: bson.D{{Key: "email_address", Value: 1}}}
	// index, _ := collection.Indexes().CreateOne(context.TODO(), indexModel)

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

	// collection.Indexes().DropOne(context.TODO(), index)

	searchQuery = bson.D{{Key: "email_address", Value: "alice@outlook.com"}}
	updatedField := bson.D{{Key: "email_address", Value: "alice@gmail.com"}}
	updater := bson.D{{Key: "$set", Value: updatedField}}
	updated, _ := collection.UpdateMany(context.TODO(), searchQuery, updater)
	if updated.ModifiedCount != 1 {
		fmt.Println("update Alice failed")
	}

	searchQuery = bson.D{{Key: "email_address", Value: "charlie@gmail.com"}}
	deleted, _ := collection.DeleteMany(context.TODO(), searchQuery)
	if deleted.DeletedCount != 1 {
		fmt.Println("delete Charlie failed")
	}

	// use Value: -1 to sort in reverse (descending) order
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: 1}})
	cursor, _ = collection.Find(context.TODO(), bson.D{}, opts)
	var documents []bson.M
	cursor.All(context.TODO(), &documents)
	for _, document := range documents {
		fmt.Println(document)
	}

	count, _ := collection.CountDocuments(context.TODO(), bson.D{})
	fmt.Println(count)
}
