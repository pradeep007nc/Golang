package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id        primitive.ObjectID `bson:"_id" gorm:"primarykey"`
	Name      string             `json:"name" gorm:"column:name"`
	Price     uint               `json:"price" gorm:"column:price"`
	CreatedAt time.Time          `json:"creted_at" gorm:"column:createdAt"`
	UpdatedAt time.Time          `json:"updated_at" gorm:"column:updatedAt"`
}
