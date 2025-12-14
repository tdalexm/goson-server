package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type CreateService struct {
	repo portsdriven.Repository
}

func NewCreateService(repo portsdriven.Repository) *CreateService {
	return &CreateService{repo}
}

func (cr *CreateService) Execute(collectionType string, record domain.Record) (domain.Record, error) {
	return cr.repo.Create(collectionType, record)
}
