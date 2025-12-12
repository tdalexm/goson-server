package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type UpdateService struct {
	repo portsdriven.Repository
}

func NewUpdateService(repo portsdriven.Repository) *UpdateService {
	return &UpdateService{repo}
}

func (s *UpdateService) Execute(collection, id string, record domain.Record) (string, error) {
	return s.repo.Update(collection, id, record)
}
