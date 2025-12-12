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

func (sr *UpdateFieldsService) Execute(collection, id string, record domain.Record) (string, error) {
	return sr.repo.UpdateFields(collection, id, record)
}
