package services

import (
	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type UpdateFieldsService struct {
	repo portsdriven.Repository
}

func NewUpdateFieldsService(repo portsdriven.Repository) *UpdateFieldsService {
	return &UpdateFieldsService{repo}
}

func (sr *UpdateFieldsService) Execute(collectionType, id string, record domain.Record) (domain.Record, error) {
	return sr.repo.UpdateFields(collectionType, id, record)
}
