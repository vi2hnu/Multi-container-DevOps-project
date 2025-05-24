package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct{
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OriginalUrl  string             `bson:"original_url" json:"original_url"`
	ShortenedUrl string             `bson:"shortened_url" json:"shortened_url"`
}