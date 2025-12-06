package services

import "github.com/tdalexm/goson-server/internal/domain"

type GetService struct {
	repo domain.Repository
}

func NewGetService(repo domain.Repository) *GetService {
	return &GetService{repo: repo}
}

func (s *GetService) Execute(resource, id string) (domain.Record, error) {
	return s.repo.GetByID(resource, id)
}
