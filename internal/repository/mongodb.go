package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository[T any] struct {
    collection *mongo.Collection
}

func NewMongoRepository[T any](db *mongo.Database , collectionName string) *MongoRepository[T] {
	return &MongoRepository[T]{collection:db.Collection(collectionName)}
}

// Inserir um novo documento
func (r *MongoRepository[T]) Insert(ctx context.Context, entity *T) (*mongo.InsertOneResult, error) {
	result, err := r.collection.InsertOne(ctx, entity)
    fmt.Println(err)
	return result, err
}

// Buscar um documento pelo ID
func (r *MongoRepository[T]) FindByID(ctx context.Context, id interface{}) (*T, error) {
	var entity T
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&entity)
	return &entity, err
}

// Listar todos os documentos
func (r *MongoRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var results []T
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item T
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		results = append(results, item)
	}
	return results, nil
}

