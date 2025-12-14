package portsdriven

import "github.com/tdalexm/goson-server/internal/domain"

type Repository interface {
	List(collectionType string) ([]domain.Record, error)
	ListWithFilter(collectionType string, filter domain.Filter) ([]domain.Record, error)
	GetByID(collectionType, id string) (domain.Record, error)
	Create(collectionType string, record domain.Record) (domain.Record, error)
	Update(collectionType, id string, record domain.Record) (domain.Record, error)
	UpdateFields(collectionType, id string, record domain.Record) (domain.Record, error)
	Delete(collectionType, id string) (string, error)
}
