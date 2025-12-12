package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type ListService struct {
	repo portsdriven.Repository
}

func NewListService(repo portsdriven.Repository) *ListService {
	return &ListService{repo: repo}
}

func (s *ListService) Execute(collection string) ([]domain.Record, error) {
	return s.repo.List(collection)
}
