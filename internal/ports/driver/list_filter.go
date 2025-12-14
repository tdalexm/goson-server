package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type ListFilterService interface {
	Execute(collectionType string, filter domain.Filter) ([]domain.Record, error)
}
