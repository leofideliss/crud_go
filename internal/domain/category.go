package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RepositoryCategory interface {
    Create(data *Category) bool
    // Read(id string) (*Category , error)
    // Update(data *Category , id string) bool
    // Delete(id string) bool
    // List() interface{}
}

type Category struct {
    Id primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name string `json:"name" bson:"name"`
    Tag string `json:"tag" bson:"tag"`
}
