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

func (s *UpdateService) Execute(collectionType, id string, record domain.Record) (domain.Record, error) {
	return s.repo.Update(collectionType, id, record)
}
