package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var (
	ErrNoMovie   = errors.New("models: no matching movie found")
	ErrDuplicate = errors.New("models: duplicate movie title")
)

type Movies struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Genre       string             `bson:"genre" json:"genre"`
	Rating      float64            `bson:"rating" json:"rating"`
	SessionTime time.Time          `bson:"sessionTime" json:"sessionTime"`
}
