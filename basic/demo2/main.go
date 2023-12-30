package main

//main funtion

import (
	"demo2/ormPrac"
	"fmt"
	"log"
)

const uri = "mongodb://localhost:27017"

func main() {
	// // Use the SetServerAPIOptions() method to set the Stable API version to 1
	// serverApi := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverApi)

	// //creating a new connection
	// client, err := mongo.Connect(context.TODO(), opts)
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// //send a ping after succesful connection
	// var result bson.M
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.M{"ping": 1}).Decode(&result); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("sucessfully connected to database")

	employees, err := ormPrac.FetchEmployees()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(employees)
}
