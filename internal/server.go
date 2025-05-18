package category

import (
	"context"

	categorypbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/category/v1"
	commonpbv1 "github.com/Ostap00034/course-work-backend-api-specs/gen/go/common/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	categorypbv1.UnimplementedCategoryServiceServer
	svc Service
}

func NewServer(s Service) *Server {
	return &Server{svc: s}
}

func (s *Server) CreateCategory(ctx context.Context, req *categorypbv1.CreateCategoryRequest) (*categorypbv1.CreateCategoryResponse, error) {
	category, err := s.svc.Create(ctx, req.Name, req.Description)
	if err != nil {
		return nil, err
	}
	return &categorypbv1.CreateCategoryResponse{
		Category: &commonpbv1.CategoryData{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.String(),
			UpdatedAt:   category.UpdatedAt.String(),
		},
	}, nil
}

func (s *Server) GetCategoryById(ctx context.Context, req *categorypbv1.GetCategoryByIdRequest) (*categorypbv1.GetCategoryByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
	}

	category, err := s.svc.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &categorypbv1.GetCategoryByIdResponse{
		Category: &commonpbv1.CategoryData{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.String(),
			UpdatedAt:   category.UpdatedAt.String(),
		},
	}, nil
}

func (s *Server) GetCategories(ctx context.Context, req *categorypbv1.GetCategoriesRequest) (*categorypbv1.GetCategoriesResponse, error) {
	categories, err := s.svc.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var categoriesList []*commonpbv1.CategoryData

	for _, category := range categories {
		categoriesList = append(categoriesList, &commonpbv1.CategoryData{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.String(),
			UpdatedAt:   category.UpdatedAt.String(),
		})
	}
	return &categorypbv1.GetCategoriesResponse{
		Categories: categoriesList,
	}, nil
}

func (s *Server) UpdateCategory(ctx context.Context, req *categorypbv1.UpdateCategoryRequest) (*categorypbv1.GetCategoryByIdResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID")
	}

	var (
		namePtr        *string
		descriptionPtr *string
	)

	if req.Category.Name != "" {
		namePtr = &req.Category.Name
	}
	if req.Category.Description != "" {
		descriptionPtr = &req.Category.Description
	}

	category, err := s.svc.Update(ctx, id, namePtr, descriptionPtr)
	if err != nil {
		return nil, err
	}

	return &categorypbv1.GetCategoryByIdResponse{
		Category: &commonpbv1.CategoryData{
			Id:          category.ID.String(),
			Name:        category.Name,
			Description: category.Description,
			CreatedAt:   category.CreatedAt.String(),
			UpdatedAt:   category.UpdatedAt.String(),
		},
	}, nil
}

func (s *Server) DeleteCategory(ctx context.Context, req *categorypbv1.DeleteCategoryRequest) (*categorypbv1.DeleteCategoryResponse, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "неправильный формат UUID")
	}
	err = s.svc.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return &categorypbv1.DeleteCategoryResponse{}, nil
}
