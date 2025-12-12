package portsdriver

import "github.com/tdalexm/goson-server/internal/domain"

type ListService interface {
	Execute(collection string) []domain.Record
}
