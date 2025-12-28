package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type GetService interface {
	Execute(collectionType, id string) (domain.Record, error)
}
