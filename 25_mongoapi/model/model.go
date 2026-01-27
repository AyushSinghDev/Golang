package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive" //Typically goes in /env but since we are only practising we used it here
)

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
