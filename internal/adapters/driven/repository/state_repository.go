package repository

import (
	"fmt"
	"strconv"

	"github.com/tdalexm/goson-server/internal/domain"
	portsdriven "github.com/tdalexm/goson-server/internal/ports/driven"
)

type StateRepository struct {
	data map[string][]domain.Record
}

func NewStateRepository(data map[string][]domain.Record) portsdriven.Repository {
	return &StateRepository{data: data}
}

func (sr *StateRepository) List(collectionType string) ([]domain.Record, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	return res, nil
}

func (sr *StateRepository) ListWithFilter(collectionType string, filters []domain.Filter) ([]domain.Record, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	var records []domain.Record
	for _, element := range res {
		matches := matchesAllFilters(element, filters)
		if matches {
			records = append(records, element)
		}
	}

	return records, nil
}

func (sr *StateRepository) GetByID(collectionType, id string) (domain.Record, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	for _, element := range res {
		if element["id"] == id {
			return element, nil
		}
	}
	return nil, domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with id '%s' not found", collectionType, id),
	)
}

func (sr *StateRepository) Create(collectionType string, record domain.Record) (domain.Record, error) {
	var res []domain.Record
	var exists bool
	if res, exists = sr.data[collectionType]; !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	if _, hasID := record["id"]; !hasID {
		newID := sr.generateNextID(res)
		record["id"] = newID
	}

	id, isStr := record["id"].(string)
	if !isStr {
		return nil, domain.NewAppError(
			domain.ErrValidation,
			"ID must be a string. Ex: 'id':'25' ",
		)
	}

	for _, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			return nil, domain.NewAppError(
				domain.ErrValidation,
				fmt.Sprintf("ID '%s' is duplicated. ID must be unique.", id),
			)
		}
	}

	sr.data[collectionType] = append(sr.data[collectionType], record)

	return record, nil
}

func (sr *StateRepository) Update(collectionType, id string, record domain.Record) (domain.Record, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	if _, hasID := record["id"]; hasID {
		return nil, domain.NewAppError(
			domain.ErrValidation,
			"Cannot update the ID field",
		)
	}

	for i, element := range res {
		elementID, ok := element["id"].(string)
		if ok && elementID == id {
			record["id"] = id
			sr.data[collectionType][i] = record
			return record, nil
		}
	}

	return nil, domain.NewAppError(
		domain.ErrCodeNotFound,
		fmt.Sprintf("%s with ID '%s' not found", collectionType, id),
	)
}

func (sr *StateRepository) UpdateFields(collectionType, id string, record domain.Record) (domain.Record, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}

	if _, hasID := record["id"]; hasID {
		return nil, domain.NewAppError(
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
		return nil, domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with ID '%s' not found", collectionType, id),
		)
	}

	current := res[foundIndex]
	for key, value := range record {
		if key != "id" {
			current[key] = value
		}
	}

	sr.data[collectionType][foundIndex] = current

	return current, nil
}

func (sr *StateRepository) Delete(collectionType, id string) (string, error) {
	res, exists := sr.data[collectionType]
	if !exists {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("collection '%s' not found in json", collectionType),
		)
	}
	foundIndex := -1
	for i, element := range res {
		if element["id"] == id {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return "", domain.NewAppError(
			domain.ErrCodeNotFound,
			fmt.Sprintf("%s with ID '%s' not found", collectionType, id),
		)
	}
	sr.data[collectionType] = append(res[:foundIndex], res[foundIndex+1:]...)
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

func matchesAllFilters(record domain.Record, filters []domain.Filter) bool {
	for _, filter := range filters {
		fieldValue, exists := record[filter.Field]
		if !exists {
			return false
		}

		matches := filter.Matches(fieldValue)
		if !matches {
			return false
		}
	}

	return true
}
