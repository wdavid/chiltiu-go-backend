package service

import (
	"Proyectos_Go/internal/core/model"
	"Proyectos_Go/internal/infrastructure/repository"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id uint) (*model.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ChangeRole(userID uint, newRole string) error {
	if newRole != "admin" && newRole != "superadmin" && newRole != "user" {
		return errors.New("rol inválido")
	}

	user, err := s.repo.GetByID(userID)
	if err != nil {
		return err
	}

	user.Role = newRole
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
