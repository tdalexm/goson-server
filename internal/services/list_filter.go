package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type ListFilterService struct {
	repo portsdriven.Repository
}

func NewListFilterService(repo portsdriven.Repository) *ListFilterService {
	return &ListFilterService{repo: repo}
}

func (s *ListFilterService) Execute(collection string, filter domain.Filter) ([]domain.Record, error) {
	return s.repo.ListWithFilter(collection, filter)
}
