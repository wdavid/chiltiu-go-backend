package service

import (
	"Proyectos_Go/internal/core/model"
	"Proyectos_Go/internal/infrastructure/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepo
}

func NewCategoryService(repo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(name, description string) error {
	cat := &model.Category{
		Name:        name,
		Description: description,
	}
	return s.repo.Create(cat)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.repo.GetAll()
}
