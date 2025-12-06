package services

import "github.com/tdalexm/goson-server/internal/domain"

type ListService struct {
	repo domain.Repository
}

func NewListService(repo domain.Repository) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Execute(resource string) ([]domain.Record, error) {
	return s.repo.List(resource)
}
