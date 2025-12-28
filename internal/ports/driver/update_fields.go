package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type UpdateFieldsService interface {
	Execute(collectionType, id string, record domain.Record) (domain.Record, error)
}
