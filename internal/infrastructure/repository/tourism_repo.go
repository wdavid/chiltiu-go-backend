package repository

import (
	"Proyectos_Go/internal/core/model"

	"gorm.io/gorm"
)

type TourismRepo struct {
	DB *gorm.DB
}

func NewTourismRepo(db *gorm.DB) *TourismRepo {
	return &TourismRepo{DB: db}
}

func (r *TourismRepo) GetAll(page, limit int) ([]model.TouristDestination, int64, error) {
	var destinations []model.TouristDestination
	var total int64

	offset := (page - 1) * limit

	if err := r.DB.Model(&model.TouristDestination{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.DB.Preload("Category").Limit(limit).Offset(offset).Find(&destinations).Error

	return destinations, total, err
}

func (r *TourismRepo) GetByID(id uint) (*model.TouristDestination, error) {
	var dest model.TouristDestination
	err := r.DB.Preload("Category").First(&dest, id).Error
	return &dest, err
}

func (r *TourismRepo) Create(dest *model.TouristDestination) error {
	return r.DB.Create(dest).Error
}

func (r *TourismRepo) Update(dest *model.TouristDestination) error {
	return r.DB.Save(dest).Error
}

func (r *TourismRepo) Delete(id uint) error {
	return r.DB.Delete(&model.TouristDestination{}, id).Error
}
