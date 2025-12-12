package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type ListFilterService interface {
	Execute(collection string, filter domain.Filter) ([]domain.Record, error)
}
