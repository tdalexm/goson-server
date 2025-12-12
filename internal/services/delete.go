package services

import portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"

type DeleteService struct {
	repo portsdriven.Repository
}

func NewDeleteService(repo portsdriven.Repository) *DeleteService {
	return &DeleteService{repo}
}

func (cr *DeleteService) Execute(collection, id string) (string, error) {
	return cr.repo.Delete(collection, id)
}
