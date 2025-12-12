package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type GetService struct {
	repo portsdriven.Repository
}

func NewGetService(repo portsdriven.Repository) *GetService {
	return &GetService{repo: repo}
}

func (s *GetService) Execute(collection, id string) (domain.Record, error) {
	return s.repo.GetByID(collection, id)
}
