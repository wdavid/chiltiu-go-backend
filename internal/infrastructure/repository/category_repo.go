package repository

import (
	"Proyectos_Go/internal/core/model"

	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) Create(cat *model.Category) error {
	return r.DB.Create(cat).Error
}

func (r *CategoryRepo) GetAll() ([]model.Category, error) {
	var categories []model.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}
