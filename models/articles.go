package models

import (
	"time"

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

// InitArticle initializes article id, created_at and updated_at fields for mongodb
func (a *Article) InitArticle() {
	t := time.Now().UTC()
	a.ID = primitive.NewObjectIDFromTimestamp(t)
	a.CreatedAt = primitive.NewDateTimeFromTime(t)
	a.UpdatedAt = primitive.NewDateTimeFromTime(t)
}

// GenerateUpdatedAtTime generates the updated timestamp for the article
func (a *Article) GenerateUpdatedAtTime() {
	a.UpdatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())
}
