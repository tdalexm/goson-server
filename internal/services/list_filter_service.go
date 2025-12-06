package services

import "github.com/tdalexm/goson-server/internal/domain"

type ListFilterService struct {
	repo domain.Repository
}

func NewListFilterService(repo domain.Repository) *ListFilterService {
	return &ListFilterService{repo: repo}
}

func (s *ListFilterService) Execute(resource string, filter domain.Filter) ([]domain.Record, error) {
	return s.repo.ListWithFilter(resource, filter)
}
