package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Author    string             `bson:"author" json:"author"`
	Title     string             `bson:"title" json:"title"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at" json:"updated_at"`
	Content   string             `bson:"content" json:"content"`
}
