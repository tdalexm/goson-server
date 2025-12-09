package repository

import (
	"fmt"
	"strconv"

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

func (sr *StateRepository) Create(resource string, record domain.Record) (string, error) {
	var collection []domain.Record
	var exists bool
	if collection, exists = sr.data[resource]; !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	if _, hasID := record["id"]; !hasID {
		newID := sr.generateNextID(collection)
		record["id"] = newID
	}

	id, isStr := record["id"].(string)
	if !isStr {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"ID must be a string. Ex: 'id':'25' ",
		)
	}

	for _, element := range collection {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			return "", domain.NewAppError(
				domain.ErrValidation,
				fmt.Sprintf("ID '%s' is duplicated. ID must be unique.", id),
			)
		}
	}

	sr.data[resource] = append(sr.data[resource], record)

	return id, nil
}

func (sr *StateRepository) Update(resource, id string, record domain.Record) (string, error) {
	res, exists := sr.data[resource]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	if _, hasID := record["id"]; hasID {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"Cannot update the ID field",
		)
	}

	for i, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			record["id"] = id
			sr.data[resource][i] = record
			return id, nil
		}
	}

	return "", domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with ID '%s' not found", resource, id),
	)
}

func (sr *StateRepository) UpdateFields(resource, id string, record domain.Record) (string, error) {
	res, exists := sr.data[resource]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("resource '%s' not found in json", resource),
		)
	}

	if _, hasID := record["id"]; hasID {
		return "", domain.NewAppError(
			domain.ErrValidation,
			"Cannot update the ID field",
		)
	}

	foundIndex := -1
	for i, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with ID '%s' not found", resource, id),
		)
	}

	current := res[foundIndex]
	for key, value := range record {
		if key != "id" {
			current[key] = value
		}
	}

	sr.data[resource][foundIndex] = current

	return id, nil
}

func (sr *StateRepository) generateNextID(collection []domain.Record) string {
	if len(collection) == 0 {
		return "1"
	}

	maxID := 0
	for _, record := range collection {
		if id, ok := record["id"].(int); ok {
			if id > maxID {
				maxID = id
			}
		} else if idStr, ok := record["id"].(string); ok {
			if id, err := strconv.Atoi(idStr); err == nil && id > maxID {
				maxID = id
			}
		}
	}

	nextID := maxID + 1
	return strconv.Itoa(nextID)
}
