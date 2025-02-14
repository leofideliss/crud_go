package service

import (
    "crud_go/internal/domain"
    "crud_go/internal/repository"
    "go.mongodb.org/mongo-driver/mongo"
    "context"
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

func (s *CategoryService) Deletecategory(ctx context.Context , id string) (*mongo.DeleteResult, error) {
	return s.repo.Delete(ctx,id)
}
