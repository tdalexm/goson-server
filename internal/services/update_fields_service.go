package services

import "github.com/tdalexm/goson-server/internal/domain"

type UpdateFieldsService struct {
	repo domain.Repository
}

func NewUpdateFieldsService(repo domain.Repository) *UpdateFieldsService {
	return &UpdateFieldsService{repo}
}

func (sr *UpdateFieldsService) Execute(resource, id string, record domain.Record) (string, error) {
	return sr.repo.UpdateFields(resource, id, record)
}
