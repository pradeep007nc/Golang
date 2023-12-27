package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	Id         primitive.ObjectID `bson:"_id"`
	Text       string             `bson:"text"`
	Title      string             `bson:"title"`
	Created_at time.Time          `bson:"created_at"`
	Updated_at time.Time          `bson:"updated_at"`
	Note_id    string             `bson:"note_id"`
}
