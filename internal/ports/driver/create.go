package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type CreateService interface {
	Execute(collectionType string, record domain.Record) (domain.Record, error)
}
