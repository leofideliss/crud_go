package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
    Id primitive.ObjectID `json:"-" bson:"_id,omitempty"`
    Name string `json:"name" bson:"name"`
    Tag string `json:"tag" bson:"tag"`
}
