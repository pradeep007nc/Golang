package ormPrac

import (
	"context"
	dbconnection "demo2/dbConnection"
	"demo2/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var employeeCollection, _ = dbconnection.OpenCollection(dbconnection.Client, "Employees")

func FetchEmployees() ([]models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var employees []models.Employee

	cursor, err := employeeCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("error fetching employees: %v", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var emp models.Employee
		if err := cursor.Decode(&emp); err != nil {
			return nil, fmt.Errorf("error decoding employee: %v", err)
		}
		employees = append(employees, emp)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return employees, nil
}
