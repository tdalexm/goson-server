package services

import "github.com/tdalexm/goson-server/internal/domain"

type DeleteService struct {
	repo domain.Repository
}

func NewDeleteService(repo domain.Repository) *DeleteService {
	return &DeleteService{repo}
}

func (cr *DeleteService) Execute(resource, id string) (string, error) {
	return cr.repo.Delete(resource, id)
}
