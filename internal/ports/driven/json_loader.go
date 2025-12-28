package portsdriven

import "github.com/tdalexm/goson-server/internal/domain"

type JsonRepo interface {
	Load() (map[string][]domain.Record, error)
	Save(map[string][]domain.Record) error
}
