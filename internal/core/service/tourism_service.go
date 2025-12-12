package service

import (
	"Proyectos_Go/internal/core/model"
	"Proyectos_Go/internal/infrastructure/repository"
	"Proyectos_Go/internal/utils"
	"mime/multipart"

	"github.com/microcosm-cc/bluemonday"
)

type TourismService struct {
	repo *repository.TourismRepo
}

func NewTourismService(repo *repository.TourismRepo) *TourismService {
	return &TourismService{repo: repo}
}

func (s *TourismService) GetAll(page, limit int) ([]model.TouristDestination, int64, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return s.repo.GetAll(page, limit)
}

func (s *TourismService) GetByID(id uint) (*model.TouristDestination, error) {
	return s.repo.GetByID(id)
}

func (s *TourismService) CreateDestination(dest *model.TouristDestination, imageFile *multipart.FileHeader, videoFile *multipart.FileHeader) error {
	p := bluemonday.UGCPolicy()
	dest.Description = p.Sanitize(dest.Description)

	if imageFile != nil {
		file, err := imageFile.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		url, err := utils.UploadToCloudinary(file, "img_"+imageFile.Filename)
		if err != nil {
			return err
		}
		dest.ImageURL = url
	}

	if videoFile != nil {
		file, err := videoFile.Open()
		if err != nil {
			return err
		}
		defer file.Close()
		url, err := utils.UploadToCloudinary(file, "vid_"+videoFile.Filename)
		if err != nil {
			return err
		}
		dest.VideoURL = url
	}

	return s.repo.Create(dest)
}

func (s *TourismService) UpdateDestination(id uint, updatedData *model.TouristDestination, imageFile *multipart.FileHeader, videoFile *multipart.FileHeader) (*model.TouristDestination, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existing.Name = updatedData.Name
	p := bluemonday.UGCPolicy()
	existing.Description = p.Sanitize(updatedData.Description)
	existing.Location = updatedData.Location
	existing.Latitude = updatedData.Latitude
	existing.Longitude = updatedData.Longitude
	existing.CategoryID = updatedData.CategoryID

	if imageFile != nil {
		file, _ := imageFile.Open()
		defer file.Close()
		url, _ := utils.UploadToCloudinary(file, "img_"+imageFile.Filename)
		existing.ImageURL = url
	}

	if videoFile != nil {
		file, _ := videoFile.Open()
		defer file.Close()
		url, _ := utils.UploadToCloudinary(file, "vid_"+videoFile.Filename)
		existing.VideoURL = url
	}

	err = s.repo.Update(existing)
	return existing, err
}

func (s *TourismService) DeleteDestination(id uint) error {
	return s.repo.Delete(id)
}
