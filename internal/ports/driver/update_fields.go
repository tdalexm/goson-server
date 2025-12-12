package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type UpdateFieldsService interface {
	Execute(collection, id string, record domain.Record) (string, error)
}
