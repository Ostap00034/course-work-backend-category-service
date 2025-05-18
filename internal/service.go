package category

import (
	"context"

	"github.com/Ostap00034/course-work-backend-category-service/ent"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, name, description string) (*ent.Category, error)
	Get(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	GetAll(ctx context.Context) ([]*ent.Category, error)
	Update(ctx context.Context, id uuid.UUID, name, description *string) (*ent.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Create(ctx context.Context, name, description string) (*ent.Category, error) {
	return s.repo.Create(ctx, name, description)
}

func (s *service) Get(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	return s.repo.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]*ent.Category, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) Update(ctx context.Context, id uuid.UUID, name, description *string) (*ent.Category, error) {
	return s.repo.Update(ctx, id, name, description)
}

func (s *service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
