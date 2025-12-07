package services

import "github.com/tdalexm/goson-server/internal/domain"

type CreateService struct {
	repo domain.Repository
}

func NewCreateService(repo domain.Repository) *CreateService {
	return &CreateService{repo}
}

func (cr *CreateService) Execute(resource string, record domain.Record) (string, error) {
	return cr.repo.Create(resource, record)
}
