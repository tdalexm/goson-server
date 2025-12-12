package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type GetService interface {
	Execute(collection, id string) (domain.Record, error)
}
