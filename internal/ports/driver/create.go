package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type CreateService interface {
	Execute(collection string, record domain.Record) (string, error)
}
