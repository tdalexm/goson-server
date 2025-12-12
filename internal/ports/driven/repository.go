package portsdriven

import "github.com/tdalexm/goson-server/internal/domain"

type Repository interface {
	List(collection string) ([]domain.Record, error)
	ListWithFilter(collection string, filter domain.Filter) ([]domain.Record, error)
	GetByID(collection, id string) (domain.Record, error)
	Create(collection string, record domain.Record) (string, error)
	Update(collection, id string, record domain.Record) (string, error)
	UpdateFields(collection, id string, record domain.Record) (string, error)
	Delete(collection, id string) (string, error)
}
