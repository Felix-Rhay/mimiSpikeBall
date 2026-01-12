package services

import (
	"mimiSpikeBall/internal/models"
	"mimiSpikeBall/internal/repository"
)

type CompteService struct {
	repo *repository.CompteRepository
}

func NewCompteService(repo *repository.CompteRepository) *CompteService {
	return &CompteService{repo: repo}
}

func (s *CompteService) ObtenirComptes() (*[]models.Compte, error) {
	return s.repo.ObtenirComptes()
}

func (s *CompteService) LoginUser(entrant *models.LoginUserRequest) (bool, error) {
	return s.repo.LoginUser(entrant)
}
