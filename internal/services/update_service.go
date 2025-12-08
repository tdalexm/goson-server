package services

import "github.com/tdalexm/goson-server/internal/domain"

type UpdateService struct {
	repo domain.Repository
}

func NewUpdateService(repo domain.Repository) *UpdateService {
	return &UpdateService{repo}
}

func (s *UpdateService) Execute(resource, id string, record domain.Record) (string, error) {
	return s.repo.Update(resource, id, record)
}
