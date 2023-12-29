package models

type Employee struct {
	Id              string `bson:"_id"`
	Employee_name   string `bson:"employeeName"`
	PhoneNo         string `json:"phone"`
	Email           string `json:"email"`
	ProfileImageUrl string `json:"profileImageUrl"`
}
