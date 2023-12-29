package dbconnection

import (
	"context"
	"demo2/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	uri        = "mongodb://localhost:27017"
	dbName     = "GlobalGroupware"
	collection = "Employees"
)

func DbConnect() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//creating a new connection
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database(dbName).Collection(collection)

	// Fetch all documents
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var employees []models.Employee
	for cursor.Next(ctx) {
		var emp models.Employee
		err := cursor.Decode(&emp)
		if err != nil {
			log.Fatal(err)
		}
		employees = append(employees, emp)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Print or process the retrieved documents
	for _, emp := range employees {
		fmt.Printf("ID: %s, Name: %s, Phone: %s, Email: %s, ProfileImageURL: %s\n", emp.Id, emp.Employee_name, emp.PhoneNo, emp.Email, emp.ProfileImageUrl)
	}

}
