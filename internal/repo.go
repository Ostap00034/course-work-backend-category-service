package category

import (
	"context"
	"errors"

	"github.com/Ostap00034/course-work-backend-category-service/ent"
	"github.com/google/uuid"
)

var (
	ErrCategoryExists       = errors.New("категория с таким названием уже существует")
	ErrCategoryNotFound     = errors.New("категория не найдена")
	ErrCreateCategoryFailed = errors.New("ошибка при создании категории")
	ErrGetCategoryFailed    = errors.New("ошибка при получении категории")
	ErrGetCategoriesFailed  = errors.New("ошибка при получении категорий")
)

type Repository interface {
	Create(ctx context.Context, name, description string) (*ent.Category, error)
	Get(ctx context.Context, id uuid.UUID) (*ent.Category, error)
	GetAll(ctx context.Context) ([]*ent.Category, error)
	Update(ctx context.Context, id uuid.UUID, name, description *string) (*ent.Category, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repo struct {
	client *ent.Client
}

func NewRepo(client *ent.Client) Repository {
	return &repo{client: client}
}

func (r *repo) Create(ctx context.Context, name, description string) (*ent.Category, error) {
	category, err := r.client.Category.Create().SetName(name).SetDescription(description).Save(ctx)
	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, ErrCategoryExists
		}
		return nil, ErrCreateCategoryFailed
	}

	return category, nil
}

func (r *repo) Get(ctx context.Context, id uuid.UUID) (*ent.Category, error) {
	category, err := r.client.Category.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrCategoryNotFound
		}
		return nil, ErrGetCategoryFailed
	}

	return category, nil
}

func (r *repo) GetAll(ctx context.Context) ([]*ent.Category, error) {
	categories, err := r.client.Category.Query().All(ctx)
	if err != nil {
		return nil, ErrGetCategoriesFailed
	}

	return categories, nil
}

func (r *repo) Update(ctx context.Context, id uuid.UUID, name, description *string) (*ent.Category, error) {

	builder := r.client.Category.
		UpdateOneID(id)

	if name != nil {
		builder = builder.SetName(*name)
	}
	if description != nil {
		builder = builder.SetDescription(*description)
	}

	updated, err := builder.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrCategoryNotFound
		}
		return nil, err
	}
	return updated, nil
}

func (r *repo) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.client.Category.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrCategoryNotFound
		}
		return err
	}
	return nil
}
