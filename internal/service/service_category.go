package service

import (
    "crud_go/internal/domain"
    "crud_go/internal/repository"
    "go.mongodb.org/mongo-driver/mongo"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"

)


type CategoryService struct {
	repo *repository.MongoRepository[domain.Category]
}

func NewCategoryService(db *mongo.Database) *CategoryService {
	return &CategoryService{
		repo: repository.NewMongoRepository[domain.Category](db, "categories"),
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, category *domain.Category) (*mongo.InsertOneResult, error) {
	return s.repo.Insert(ctx, category)
}

func (s *CategoryService) ListAllCategories(ctx context.Context) ([]domain.Category, error) {
	return s.repo.FindAll(ctx)
}

func (s *CategoryService) ReadCategories(ctx context.Context , id primitive.ObjectID) (*domain.Category, error) {
	return s.repo.FindByID(ctx , id)
}

func (s *CategoryService) DeleteCategory(ctx context.Context , id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.repo.Delete(ctx,id)
}

func (s *CategoryService) ReadCategory(ctx context.Context , id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return s.repo.Delete(ctx,id)
}

func (s *CategoryService) UpdateCategory(ctx context.Context , category *domain.Category , id primitive.ObjectID) (*mongo.UpdateResult, error) {
	return s.repo.Update(ctx,category,id)
}
