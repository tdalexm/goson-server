package jsonloader

import (
	"encoding/json"
	"os"

	"github.com/tdalexm/goson-server/internal/domain"
	ports "github.com/tdalexm/goson-server/internal/ports/driven"
)

type Repo struct {
	path string
}

func NewJsonRepo(path string) ports.JsonRepo {
	return &Repo{path: path}
}

func (r *Repo) Load() (map[string][]domain.Record, error) {
	data, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}
	var db map[string][]domain.Record
	err = json.Unmarshal(data, &db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *Repo) Save(map[string][]domain.Record) error {
	return nil
}
