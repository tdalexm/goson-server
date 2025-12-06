package repository

import (
	"fmt"

	"github.com/tdalexm/goson-server/internal/domain"
)

type StateRepository struct {
	data map[string][]domain.Record
}

func NewStateRepository(data map[string][]domain.Record) domain.Repository {
	return &StateRepository{data: data}
}

func (sr *StateRepository) List(resource string) ([]domain.Record, error) {
	res, exists := sr.data[resource]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	return res, nil
}

func (sr *StateRepository) ListWithFilter(resource string, filter domain.Filter) ([]domain.Record, error) {
	res, exists := sr.data[resource]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	var records []domain.Record
	for _, element := range res {
		fieldValue, exist := element[filter.Field]
		if !exist {
			continue
		}

		match, err := filter.Matches(fieldValue)
		if err != nil {
			return nil, err
		}

		if match {
			records = append(records, element)
		}
	}

	if len(records) == 0 {
		searchValue := filter.Value
		if filter.Contains != "" {
			searchValue = filter.Contains
		}
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with %s '%s' not found", resource, filter.Field, searchValue),
		)
	}

	return records, nil
}

func (sr *StateRepository) GetByID(resource, id string) (domain.Record, error) {
	res, exists := sr.data[resource]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	for _, element := range res {
		if element["id"] == id {
			return element, nil
		}
	}
	return nil, domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with id '%s' not found", resource, id),
	)
}
