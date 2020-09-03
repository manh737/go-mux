package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Image struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	ImageLink string             `bson:"image_link"`
}

func NewImage(createdAt, updatedAt time.Time, imageLink string) *Image {
	return &Image{
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		ImageLink: imageLink,
	}
}
